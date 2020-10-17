package radiusd

import (
	"time"

	"layeh.com/radius"

	"github.com/ca17/teamsacs/constant"
	"github.com/ca17/teamsacs/models"
	"github.com/ca17/teamsacs/radiusd/radlog"
	"github.com/ca17/teamsacs/radiusd/radparser"
)


func (s *AcctService) DoAcctUpdateBefore(r *radius.Request, vr *radparser.VendorRequest,  user *models.Subscribe, vpe *models.Vpe, nasrip string) {
	// 用户状态变更为停用后触发下线
	if user.Status == constant.DISABLED {
		s.DoAcctDisconnect(r, vpe, user.Username, nasrip)
	}

	// 用户过期后触发下线
	if user.ExpireTime.Time().Before(time.Now()) {
		s.DoAcctDisconnect(r, vpe, user.Username, nasrip)
	}

	s.DoAcctUpdate(r, vr, user.Username, vpe, nasrip)
}

func (s *AcctService) DoAcctUpdate(r *radius.Request, vr *radparser.VendorRequest,  username string, vpe *models.Vpe, nasrip string) {
	online := GetRadiusOnlineFromRequest(r, vr, vpe, nasrip)
	// 更新在线信息
	err := s.Manager.GetRadiusManager().UpdateRadiusOnline(online)
	if err != nil {
		radlog.Errorf("UpdateRadiusOnlineData user:%s error, %s", username, err.Error())
	}

}
