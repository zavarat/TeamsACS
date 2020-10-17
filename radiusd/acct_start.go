package radiusd

import (
	"layeh.com/radius"

	"github.com/ca17/teamsacs/models"
	"github.com/ca17/teamsacs/radiusd/radlog"
	"github.com/ca17/teamsacs/radiusd/radparser"
)

func (s *AcctService) DoAcctStart(r *radius.Request, vr *radparser.VendorRequest,  username string, vpe *models.Vpe, nasrip string) {
	online := GetRadiusOnlineFromRequest(r, vr, vpe, nasrip)
	err := s.Manager.GetRadiusManager().AddRadiusOnline(online)
	if err!= nil {
		radlog.Errorf("AddRadiusOnline user:%s error %s", username, err.Error())
	}
}
