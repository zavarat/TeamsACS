package radiusd

import (
	"crypto/tls"
	"fmt"
	"strconv"
	"strings"
	"time"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"

	"github.com/go-ldap/ldap/v3"

	"github.com/ca17/teamsacs/common/aes"
	"github.com/ca17/teamsacs/common/mfa"
	"github.com/ca17/teamsacs/common/timeutil"
	"github.com/ca17/teamsacs/constant"
	"github.com/ca17/teamsacs/models"
	"github.com/ca17/teamsacs/radiusd/radparser"
	"github.com/ca17/teamsacs/radiusd/vendors/microsoft"
)

type LdapRadisProfile struct {
	AuthorizationProfile
	Status    string
	MacAddr   string
	MfaSecret string
	MfaStatus string
	ActiveNum int
}

//goland:noinspection ALL
func (s *AuthService) LdapUserAuth(
	rw radius.ResponseWriter,
	r *radius.Request,
	username string,
	ldapNode *models.Ldap,
	radAccept *radius.Packet,
	vreq *radparser.VendorRequest, ) (*LdapRadisProfile, error) {

	ignoreChk := s.GetStringConfig(constant.RadiusIgnorePwd, constant.DISABLED) == constant.ENABLED

	var checkType = "pap"
	// mschapv2
	challenge := microsoft.MSCHAPChallenge_Get(r.Packet)
	if challenge != nil {
		checkType = "mschapv2"
	}

	// chap
	chapPassword := rfc2865.CHAPPassword_Get(r.Packet)
	if chapPassword != nil {
		checkType = "chap"
	}

	// connect ldap
	ld, err := ldap.Dial("tcp", ldapNode.Address)
	if err != nil {
		return nil, fmt.Errorf("Username:%s ldap auth error, ldap connect error", username)
	}
	defer ld.Close()

	// start tls
	if ldapNode.Istls == constant.ENABLED {
		err = ld.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return nil, fmt.Errorf("Username:%s ldap auth error, ldap tls error", username)
		}
	}

	ldapPwd, err := aes.DecryptFromB64(ldapNode.Password, s.GetAppConfig().System.Aeskey)
	if err != nil {
		return nil, fmt.Errorf("Username:%s ldap auth error, ldap:%s password format error", username, ldapNode.Name)
	}

	err = ld.Bind(ldapNode.Basedn, ldapPwd)
	if err != nil {
		return nil, fmt.Errorf("Username:%s ldap auth error, ldap:%s %s bind auth error", username, ldapNode.Name, ldapNode.Basedn)
	}

	searchRequest := ldap.NewSearchRequest(
		ldapNode.Searchdn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(ldapNode.UserFilter, username),
		[]string{"dn", "radiusReplyItem", "radiusCallingStationId"},
		nil,
	)

	sr, err := ld.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("Username:%s ldap auth error, ldap:%s search error", username, ldapNode.Name)
	}

	if len(sr.Entries) == 0 && !ignoreChk {
		return nil, fmt.Errorf("Username:%s ldap auth error, user not exists", username)
	}

	// parse ldap radius attr
	var userProfile = new(LdapRadisProfile)
	userProfile.ExpireTime = time.Now().Add(time.Hour * 24)
	userProfile.InterimInterval = int(s.GetIntConfig(constant.AcctInterimInterval, 120))
	parseLdapRadiusAttrs(sr.Entries[0].GetAttributeValues("radiusReplyItem"), userProfile)
	userProfile.MacAddr = sr.Entries[0].GetAttributeValue("radiusCallingStationId")

	// check status
	if userProfile.Status == constant.DISABLED {
		return nil, fmt.Errorf("user:%s Ldap is disabled", username)
	}

	// check expire
	if userProfile.ExpireTime.Before(time.Now()) {
		return nil, fmt.Errorf("user:%s Ldap is expire", username)
	}

	// mac auth check
	if vreq.Macaddr == username {
		return userProfile, nil
	}

	isOtp := s.GetStringConfig(constant.RadiusMfaStatus, constant.DISABLED) == constant.ENABLED

	// 如果是 PAP 验证并且不是 OTP 验证， 直接校验 Ldap 密码
	if !ignoreChk && checkType == "pap" && !isOtp {
		password := rfc2865.UserPassword_GetString(r.Packet)
		userdn := sr.Entries[0].DN
		err = ld.Bind(userdn, password)
		if err != nil {
			return nil, fmt.Errorf("Username:%s ldap auth error, user password check error", username)
		}
	}

	// ldap otp check
	if userProfile.MfaSecret != "" && userProfile.MfaStatus == constant.ENABLED && isOtp {
		lpwd, err := mfa.NewGoogleAuth().GetCode(userProfile.MfaSecret)
		if err != nil {
			return nil, fmt.Errorf("user:%s Ldap OTP password is invalid", username)
		}
		return userProfile, s.CheckPassword(r, username, lpwd, radAccept, false)
	}

	if !ignoreChk && checkType == "chap" && !isOtp {
		return nil, fmt.Errorf("user:%s Ldap chap password is not support", username)
	}

	// check online
	err = s.CheckOnlineCount(username, userProfile.ActiveNum)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}

func parseLdapRadiusAttrs(values []string, p *LdapRadisProfile) {
	for _, value := range values {
		kv := strings.Split(value, "=")
		if len(kv) != 2 {
			continue
		}
		switch strings.TrimSpace(kv[0]) {
		case "Status":
			p.Status = strings.TrimSpace(kv[1])
		case "MfaSecret":
			p.MfaSecret = strings.TrimSpace(kv[1])
		case "MfaStatus":
			p.MfaStatus = strings.TrimSpace(kv[1])
		case "Domain":
			p.Domain = strings.TrimSpace(kv[1])
		case "AddrPool":
			p.AddrPool = strings.TrimSpace(kv[1])
		case "Ipaddr":
			p.Ipaddr = strings.TrimSpace(kv[1])
		case "LimitPolicy":
			p.LimitPolicy = strings.TrimSpace(kv[1])
		case "UpLimitPolicy":
			p.UpLimitPolicy = strings.TrimSpace(kv[1])
		case "DownLimitPolicy":
			p.DownLimitPolicy = strings.TrimSpace(kv[1])
		case "ActiveNum":
			_ActiveNum, err := strconv.ParseInt(kv[1], 10, 64)
			if err == nil {
				p.ActiveNum = int(_ActiveNum)
			}
		case "UpRateKbps":
			_UpRate, err := strconv.ParseInt(kv[1], 10, 64)
			if err == nil {
				p.UpRateKbps = int(_UpRate)
			}
		case "DownRateKbps":
			_DownRate, err := strconv.ParseInt(kv[1], 10, 64)
			if err == nil {
				p.DownRateKbps = int(_DownRate)
			}
		case "ExpireTime":
			if kv[1] != "" {
				_ExpireTime, err := time.Parse(timeutil.YYYYMMDD_LAYOUT, kv[1])
				if err == nil {
					p.ExpireTime = _ExpireTime
				}
			}
		}
	}
}
