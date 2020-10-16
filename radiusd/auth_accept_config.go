package radiusd

import (
	"fmt"
	"math"
	"net"
	"time"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"layeh.com/radius/rfc2869"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/constant"
	"github.com/ca17/teamsacs/models"
	"github.com/ca17/teamsacs/radiusd/vendors/cisco"
	"github.com/ca17/teamsacs/radiusd/vendors/h3c"
	"github.com/ca17/teamsacs/radiusd/vendors/huawei"
	"github.com/ca17/teamsacs/radiusd/vendors/ikuai"
	"github.com/ca17/teamsacs/radiusd/vendors/mikrotik"
	"github.com/ca17/teamsacs/radiusd/vendors/radback"
	"github.com/ca17/teamsacs/radiusd/vendors/zte"
)

// 用户属性策略下发配置
func (s *AuthService) AcceptAcceptConfig(user *models.Subscribe, vendorCode string, radAccept *radius.Packet) {
	configDefaultAccept(s, user, radAccept)
	switch vendorCode {
	case VendorHuawei:
		configHuaweiAccept(user, radAccept)
	case VendorH3c:
		configH3cAccept(user, radAccept)
	case VendorRadback:
		configRadbackAccept(user, radAccept)
	case VendorZte:
		configZteAccept(user, radAccept)
	case VendorCisco:
		configCiscoAccept(user, radAccept)
	case VendorMikrotik:
		configMikroTikAccept(user, radAccept)
	case VendorIkuai:
		configIkuaiAccept(user, radAccept)
	}
}

// 设置标准 RADIUS 属性
func configDefaultAccept(s *AuthService, user *models.Subscribe, radAccept *radius.Packet) {
	var timeout = int64(user.ExpireTime.Time().Sub(time.Now()).Seconds())
	if timeout > math.MaxInt32 {
		timeout = math.MaxInt32
	}
	var interimTimes = s.GetIntConfig(constant.AcctInterimInterval, 120)
	rfc2865.SessionTimeout_Set(radAccept, rfc2865.SessionTimeout(timeout))
	rfc2869.AcctInterimInterval_Set(radAccept, rfc2869.AcctInterimInterval(interimTimes))

	if common.IsNotEmptyAndNA(user.Profile.AddrPool) {
		rfc2869.FramedPool_SetString(radAccept, user.Profile.AddrPool)
	}

	if common.IsNotEmptyAndNA(user.Ipaddr) {
		rfc2865.FramedIPAddress_Set(radAccept, net.ParseIP(user.Ipaddr))
	}
}

func configMikroTikAccept(user *models.Subscribe, radAccept *radius.Packet) {
	mikrotik.MikrotikRateLimit_SetString(radAccept, fmt.Sprintf("%dk/%dk", user.Profile.UpRate, user.Profile.DownRate))
}

func configIkuaiAccept(user *models.Subscribe, radAccept *radius.Packet) {
	var up = int64(user.Profile.UpRate) * 1024 * 8
	var down = int64(user.Profile.DownRate) * 1024 * 8
	if up > math.MaxInt32 {
		up = math.MaxInt32
	}
	if down > math.MaxInt32 {
		down = math.MaxInt32
	}

	ikuai.RPUpstreamSpeedLimit_Set(radAccept, ikuai.RPUpstreamSpeedLimit(up))
	ikuai.RPDownstreamSpeedLimit_Set(radAccept, ikuai.RPDownstreamSpeedLimit(down))
}

func configHuaweiAccept(user *models.Subscribe, radAccept *radius.Packet) {
	var up = int64(user.Profile.UpRate) * 1024
	var down = int64(user.Profile.DownRate) * 1024
	var upPeak = up * 4
	var downPeak = down * 4
	if up > math.MaxInt32 {
		up = math.MaxInt32
	}
	if upPeak > math.MaxInt32 {
		upPeak = math.MaxInt32
	}
	if down > math.MaxInt32 {
		down = math.MaxInt32
	}
	if downPeak > math.MaxInt32 {
		downPeak = math.MaxInt32
	}
	huawei.HuaweiInputAverageRate_Set(radAccept, huawei.HuaweiInputAverageRate(up))
	huawei.HuaweiInputPeakRate_Set(radAccept, huawei.HuaweiInputPeakRate(upPeak))
	huawei.HuaweiOutputAverageRate_Set(radAccept, huawei.HuaweiOutputAverageRate(down))
	huawei.HuaweiOutputPeakRate_Set(radAccept, huawei.HuaweiOutputPeakRate(downPeak))

	if common.IsNotEmptyAndNA(user.Profile.Domain) {
		huawei.HuaweiDomainName_SetString(radAccept, user.Profile.Domain)
	}
}

func configH3cAccept(user *models.Subscribe, radAccept *radius.Packet) {
	var up = int64(user.Profile.UpRate) * 1024
	var down = int64(user.Profile.DownRate) * 1024
	var upPeak = up * 4
	var downPeak = down * 4
	if up > math.MaxInt32 {
		up = math.MaxInt32
	}
	if upPeak > math.MaxInt32 {
		upPeak = math.MaxInt32
	}
	if down > math.MaxInt32 {
		down = math.MaxInt32
	}
	if downPeak > math.MaxInt32 {
		downPeak = math.MaxInt32
	}
	h3c.H3CInputAverageRate_Set(radAccept, h3c.H3CInputAverageRate(up))
	h3c.H3CInputPeakRate_Set(radAccept, h3c.H3CInputPeakRate(upPeak))
	h3c.H3COutputAverageRate_Set(radAccept, h3c.H3COutputAverageRate(down))
	h3c.H3COutputPeakRate_Set(radAccept, h3c.H3COutputPeakRate(downPeak))
}

func configRadbackAccept(user *models.Subscribe, radAccept *radius.Packet) {
	if common.IsNotEmptyAndNA(user.Profile.LimitPolicy) {
		radback.SubscriberProfileName_SetString(radAccept, user.Profile.LimitPolicy)
	}
	if common.IsNotEmptyAndNA(user.Profile.Domain) {
		radback.ContextName_SetString(radAccept, user.Profile.Domain)
	}
}

func configZteAccept(user *models.Subscribe, radAccept *radius.Packet) {
	var up = int64(user.Profile.UpRate) * 1024
	var down = int64(user.Profile.DownRate) * 1024
	if up > math.MaxInt32 {
		up = math.MaxInt32
	}
	if down > math.MaxInt32 {
		down = math.MaxInt32
	}
	zte.ZTERateCtrlSCRUp_Set(radAccept, zte.ZTERateCtrlSCRUp(up))
	zte.ZTERateCtrlSCRDown_Set(radAccept, zte.ZTERateCtrlSCRDown(down))
	if common.IsNotEmptyAndNA(user.Profile.Domain) {
		zte.ZTEContextName_SetString(radAccept, user.Profile.Domain)
	}
}

func configCiscoAccept(user *models.Subscribe, radAccept *radius.Packet) {
	if common.IsNotEmptyAndNA(user.Profile.UpLimitPolicy) {
		cisco.CiscoAVPair_Add(radAccept, []byte(fmt.Sprintf("sub-qos-policy-in=%s", user.Profile.UpLimitPolicy)))
	}
	if common.IsNotEmptyAndNA(user.Profile.DownLimitPolicy) {
		cisco.CiscoAVPair_Add(radAccept, []byte(fmt.Sprintf("sub-qos-policy-out=%s", user.Profile.DownLimitPolicy)))
	}
}
