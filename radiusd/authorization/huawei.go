package authorization

import (
	"math"

	"layeh.com/radius"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/radiusd/vendors/huawei"
)

func HuaweiAuthorization(prof Profile, accept *radius.Packet) {
	var up = prof.GetUpRateKbps() * 1000
	var down = prof.GetDownRateKbps() * 1000
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
	huawei.HuaweiInputAverageRate_Set(accept, huawei.HuaweiInputAverageRate(up))
	huawei.HuaweiInputPeakRate_Set(accept, huawei.HuaweiInputPeakRate(upPeak))
	huawei.HuaweiOutputAverageRate_Set(accept, huawei.HuaweiOutputAverageRate(down))
	huawei.HuaweiOutputPeakRate_Set(accept, huawei.HuaweiOutputPeakRate(downPeak))

	domain := prof.GetDomain()
	if common.IsNotEmptyAndNA(domain) {
		huawei.HuaweiDomainName_SetString(accept, domain)
	}
}
