package authorization

import (
	"layeh.com/radius"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/radiusd/vendors/radback"
)

func RadbackAuthorization(prof Profile, accept *radius.Packet) {
	limitPolicy := prof.GetLimitPolicy()
	if common.IsNotEmptyAndNA(limitPolicy) {
		radback.SubscriberProfileName_SetString(accept, limitPolicy)
	}
	domain := prof.GetDomain()
	if common.IsNotEmptyAndNA(domain) {
		radback.ContextName_SetString(accept, domain)
	}
}
