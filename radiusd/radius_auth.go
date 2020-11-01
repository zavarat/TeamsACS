package radiusd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"

	"github.com/ca17/teamsacs/radiusd/authorization"
	"github.com/ca17/teamsacs/radiusd/debug"
	"github.com/ca17/teamsacs/radiusd/radlog"
	"github.com/ca17/teamsacs/radiusd/radparser"
)

// 认证服务
type AuthService struct {
	*RadiusService
}

func NewAuthService(radiusService *RadiusService) *AuthService {
	return &AuthService{RadiusService: radiusService}
}

// RADIUS Auth
func (s *AuthService) ServeRADIUS(w radius.ResponseWriter, r *radius.Request) {
	var start = time.Now()
	defer func() {
		if ret := recover(); ret != nil {
			err, ok := ret.(error)
			if ok {
				radlog.Error(err)
				s.SendReject(w, r, err.Error())
			}
		}
	}()

	if s.GetAppConfig().Radiusd.Debug {
		radlog.Info(debug.FmtRequest(r))
	}

	// nas access check
	raddrstr := r.RemoteAddr.String()
	ip := raddrstr[:strings.Index(raddrstr, ":")]
	var identifier = rfc2865.NASIdentifier_GetString(r.Packet)
	username := rfc2865.UserName_GetString(r.Packet)

	// Username empty  check
	if username == "" {
		s.CheckRadAuthError(start, rfc2865.CallingStationID_GetString(r.Packet), ip, errors.New("username is empty of client mac"))
	}

	vpe, err := s.GetNas(ip, identifier)
	s.CheckRadAuthError(start, username, ip, err)

	//  setup new packet secret
	r.Secret = []byte(vpe.Secret)
	r.Packet.Secret = []byte(vpe.Secret)
	response := r.Response(radius.CodeAccessAccept)

	vendorReq := radparser.ParseVendor(r, vpe.VendorCode)

	// ----------------------------------------------------------------------------------------------------
	// Ldap auth
	if vpe.LdapId == ""{
		lnode, err := s.GetLdap(vpe.LdapId)
		s.CheckRadAuthError(start, username, ip, err)
		// check ldap auth
		userProfile, err := s.LdapUserAuth(w, r, username, lnode, response, vendorReq)
		s.CheckRadAuthError(start, username, ip, err)
		// setup accept
		authorization.UpdateAuthorization(userProfile, vpe.VendorCode, response)
		// if ok
		s.SendAccept(w, r, response)
		s.LogAuthSucess(start, username, ip)
		return
	}

	// ----------------------------------------------------------------------------------------------------
	// Fetch validate user
	isMacAuth := vendorReq.Macaddr == username
	user, err := s.GetUser(username, isMacAuth)
	s.CheckRadAuthError(start, username, ip, err)

	if !isMacAuth {

		if user.Profile.ActiveNum != 0 {
			onlineCount, _ := s.Manager.GetRadiusManager().GetOnlineCount(user.Username)
			if int(onlineCount) > user.Profile.ActiveNum {
				s.CheckRadAuthError(start, username, ip, fmt.Errorf("user:%s active num over limit(max=%d)", user.Username, user.Profile.ActiveNum))
			}
		}

		s.CheckRadAuthError(start, username, ip, s.CheckOnlineCount(username, user.Profile.ActiveNum))

		// Username Mac bind check
		s.CheckRadAuthError(start, username, ip, s.CheckMacBind(user, vendorReq))

		// Username vlanid check
		s.CheckRadAuthError(start, username, ip, s.CheckVlanBind(user, vendorReq))
	}

	// Password check
	// if mschapv2 auth, will set accept attribute
	localpwd, err := s.GetLocalPassword(user, isMacAuth)
	s.CheckRadAuthError(start, username, ip, err)
	s.CheckRadAuthError(start, username, ip, s.CheckPassword(r, user.Username, localpwd, response, isMacAuth))

	// setup accept
	authorization.UpdateAuthorization(user, vpe.VendorCode, response)

	// send accept
	s.SendAccept(w, r, response)
	// update mac & vlan
	s.UpdateBind(user, vendorReq)

	s.LogAuthSucess(start, username, ip)
}

// send accept
func (s *AuthService) SendAccept(w radius.ResponseWriter, r *radius.Request, resp *radius.Packet) {
	radlog.Infof("Writing %v to %v", resp.Code, r.RemoteAddr)
	if s.GetAppConfig().Radiusd.Debug {
		radlog.Info(debug.FmtResponse(resp, r.RemoteAddr))
	}
	err := w.Write(resp)
	if err != nil {
		radlog.Error(err)
	}
}

// send reject
func (s *AuthService) SendReject(w radius.ResponseWriter, r *radius.Request, message string) {
	defer func() {
		if ret := recover(); ret != nil {
			err, ok := ret.(error)
			if ok {
				radlog.Error(err)
			}
		}
	}()

	var code = radius.CodeAccessReject
	var resp = r.Response(code)
	if message != "" {
		if len(message) > 253 {
			message = message[0:252]
		}
		_ = rfc2865.ReplyMessage_SetString(resp, message)
	}
	radlog.Infof("Writing %v to %v", code, r.RemoteAddr)
	if s.GetAppConfig().Radiusd.Debug {
		radlog.Info(debug.FmtResponse(resp, r.RemoteAddr))
	}
	err := w.Write(resp)
	if err != nil {
		radlog.Error(err)
	}
}
