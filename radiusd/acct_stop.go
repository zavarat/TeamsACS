package radiusd

import (
	"layeh.com/radius"

	"github.com/ca17/teamsacs/models"
)

func (s *AcctService) DoAcctStop(r *radius.Request, vr *VendorRequest,  username string, vpe *models.Vpe, nasrip string) {
	online := GetRadiusOnlineFromRequest(r, vr, vpe, nasrip)
	if err := s.Manager.GetRadiusManager().AddRadiusAccounting(online); err!=nil {
		radlog.Errorf("AddRadiusAccounting user:%s error %s ", username, err.Error())
	}
	if err := s.Manager.GetRadiusManager().DeleteRadiusOnline(online.AcctSessionId); err != nil {
		radlog.Errorf("DeleteRadiusOnline user:%s error %s ", username, err.Error())
	}
}
