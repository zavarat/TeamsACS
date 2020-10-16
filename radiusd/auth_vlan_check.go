package radiusd

import (
	"errors"

	"github.com/ca17/teamsacs/models"
)

// vlanid  绑定检测
// 用户 vlanid 和请求 vlanid 同时有效才检测
// 如果用户 vlanid 为空, 就直接更新用户 vlanid
func (s *AuthService) CheckVlanBind(user *models.Subscribe, vendorReq *VendorRequest) error {
	reqvid := int(vendorReq.Vlanid1)
	if user.BindVlan == 1 {
		if user.Vlanid1 != 0 && vendorReq.Vlanid1 != 0 {
			if user.Vlanid1 != reqvid {
				return errors.New("user vlanid1 bind not match")
			}
		} else {
			s.UpdateUserVlanid1(user.Username, reqvid)
		}
	}

	reqvid2 := int(vendorReq.Vlanid2)
	if user.BindVlan == 1 {
		if user.Vlanid2 != 0 && reqvid2 != 0 {
			if user.Vlanid2 != reqvid2 {
				return errors.New("user vlanid2 bind not match")
			}
		} else {
			s.UpdateUserVlanid2(user.Username, reqvid2)
		}
	}

	return nil
}
