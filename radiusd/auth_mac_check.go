package radiusd

import (
	"errors"

	"github.com/ca17/teamsacs/models"
)

// mac 绑定检测
// 用户mac和请求mac同时有效才检测
// 如果用户mac为空, 就直接更新用户mac
func (s *AuthService) CheckMacBind(user *models.Subscribe, vendorReq *VendorRequest) error {
	if user.BindMac == 1 {
		if user.Macaddr != "" && user.Macaddr != "N/A" && vendorReq.Macaddr != "" {
			if user.Macaddr != vendorReq.Macaddr {
				return errors.New("user mac bind not match")
			}
		}
	}
	return nil
}
