package authorization

import (
	"math"

	"layeh.com/radius"

	"github.com/ca17/teamsacs/radiusd/vendors/h3c"
)

func H3CAuthorization(prof Profile, accept *radius.Packet) {
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
	h3c.H3CInputAverageRate_Set(accept, h3c.H3CInputAverageRate(up))
	h3c.H3CInputPeakRate_Set(accept, h3c.H3CInputPeakRate(upPeak))
	h3c.H3COutputAverageRate_Set(accept, h3c.H3COutputAverageRate(down))
	h3c.H3COutputPeakRate_Set(accept, h3c.H3COutputPeakRate(downPeak))
}
