package radiusd

import (
	"layeh.com/radius"
	"layeh.com/radius/rfc2866"

	"github.com/ca17/teamsacs/models"
	"github.com/ca17/teamsacs/radiusd/radparser"
)

func (s *AcctService) LdapUserAcct(r *radius.Request, vr *radparser.VendorRequest, username string, vpe *models.Vpe, nasrip string) {
	statusType := rfc2866.AcctStatusType_Get(r.Packet)
	switch statusType {
	case rfc2866.AcctStatusType_Value_Start:
		s.processAcctStart(r, vr, username, vpe, nasrip)
	case rfc2866.AcctStatusType_Value_InterimUpdate:
		s.processAcctUpdate(r, vr, username, vpe, nasrip)
	case rfc2866.AcctStatusType_Value_Stop:
		s.processAcctStop(r, vr, username, vpe, nasrip)
	case rfc2866.AcctStatusType_Value_AccountingOn:
		s.processAcctNasOn(r)
	case rfc2866.AcctStatusType_Value_AccountingOff:
		s.processAcctNasOff(r)
	}
}
