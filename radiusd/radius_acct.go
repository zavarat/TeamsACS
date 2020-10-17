package radiusd

import (
	"errors"
	"strings"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"layeh.com/radius/rfc2866"

	"github.com/ca17/teamsacs/radiusd/radlog"
	"github.com/ca17/teamsacs/radiusd/radparser"
)

// 记账服务
type AcctService struct {
	*RadiusService
}

func NewAcctService(radiusService *RadiusService) *AcctService {
	return &AcctService{RadiusService: radiusService}
}

func (s *AcctService) ServeRADIUS(w radius.ResponseWriter, r *radius.Request) {
	defer func() {
		if ret := recover(); ret != nil {
			err, ok := ret.(error)
			if ok {
				radlog.Error(err)
			}
		}
	}()

	if s.GetAppConfig().Radiusd.Debug {
		radlog.Info(FmtRequest(r))
	}

	// NAS 接入检查
	raddrstr := r.RemoteAddr.String()
	nasrip := raddrstr[:strings.Index(raddrstr, ":")]
	var identifier = rfc2865.NASIdentifier_GetString(r.Packet)
	vpe, err := s.GetNas(nasrip, identifier)
	radlog.CheckError(err)

	// 重新设置数据报文秘钥
	r.Secret = []byte(vpe.Secret)
	r.Packet.Secret = []byte(vpe.Secret)

	// 用户名检查
	username := rfc2865.UserName_GetString(r.Packet)
	if username == "" {
		radlog.CheckError(errors.New("username is empty"))
	}

	vendorReq := radparser.ParseVendor(r, vpe.VendorCode)

	// Ldap acct
	if vpe.LdapId.IsZero() {
		_, err := s.GetLdap(vpe.LdapId)
		radlog.CheckError(err)
		// check ldap auth
		s.LdapUserAcct(r, vendorReq, username, vpe, nasrip)
		// if ok
		s.SendResponse(w, r)
		return
	}

	// 获取有效用户
	user, err := s.GetUserForAcct(username)
	radlog.CheckError(err)

	statusType := rfc2866.AcctStatusType_Get(r.Packet)
	switch statusType {
	case rfc2866.AcctStatusType_Value_Start:
		s.DoAcctStart(r, vendorReq, user.Username, vpe, nasrip)
	case rfc2866.AcctStatusType_Value_InterimUpdate:
		s.DoAcctUpdateBefore(r, vendorReq, user, vpe, nasrip)
	case rfc2866.AcctStatusType_Value_Stop:
		s.DoAcctStop(r, vendorReq, user.Username, vpe, nasrip)
	case rfc2866.AcctStatusType_Value_AccountingOn:
		s.DoAcctNasOn(r)
	case rfc2866.AcctStatusType_Value_AccountingOff:
		s.DoAcctNasOff(r)
	}

	s.SendResponse(w, r)
}

func (s *AcctService) SendResponse(w radius.ResponseWriter, r *radius.Request) {
	resp := r.Response(radius.CodeAccountingResponse)
	err := w.Write(resp)
	radlog.Infof("Writing %v to %v", resp.Code, r.RemoteAddr)
	if s.GetAppConfig().Radiusd.Debug {
		radlog.Info(FmtResponse(resp, r.RemoteAddr))
	}
	if err != nil {
		radlog.Error(err)
	}
}
