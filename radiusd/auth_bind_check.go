package radiusd

import (
	"errors"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/models"
	"github.com/ca17/teamsacs/radiusd/radparser"
)

// CheckVlanBind
// vlanid binding detection
// Only if both user vlanid and request vlanid are valid.
// If user vlanid is empty, update user vlanid directly.
func (s *AuthService) CheckVlanBind(user *models.Subscribe, vendorReq *radparser.VendorRequest) error {
	if user.BindVlan == 0 {
		return nil
	}
	reqvid1 := int(vendorReq.Vlanid1)
	reqvid2 := int(vendorReq.Vlanid2)
	if user.Vlanid1 != 0 && vendorReq.Vlanid1 != 0 && user.Vlanid1 != reqvid1 {
		return errors.New("user vlanid1 bind not match")
	}

	if user.Vlanid2 != 0 && reqvid2 != 0 && user.Vlanid2 != reqvid2 {
		return errors.New("user vlanid2 bind not match")
	}

	return nil
}

// CheckMacBind
// mac binding detection
// Detected only if both user mac and request mac are valid.
// If user mac is empty, update user mac directly.
func (s *AuthService) CheckMacBind(user *models.Subscribe, vendorReq *radparser.VendorRequest) error {
	if user.BindMac == 0 {
		return nil
	}

	if common.IsNotEmptyAndNA(user.Macaddr) && vendorReq.Macaddr != "" && user.Macaddr != vendorReq.Macaddr {
		return errors.New("user mac bind not match")
	}
	return nil
}


// UpdateBind
// update mac or vlan
func (s *AuthService) UpdateBind(user *models.Subscribe, vendorReq *radparser.VendorRequest) {
	if user.Macaddr != vendorReq.Macaddr {
		s.UpdateUserMac(user.Username, vendorReq.Macaddr)
	}
	reqvid1 := int(vendorReq.Vlanid1)
	reqvid2 := int(vendorReq.Vlanid2)
	if user.Vlanid1 != reqvid1 {
		s.UpdateUserVlanid2(user.Username, reqvid1)
	}
	if user.Vlanid2 != reqvid2 {
		s.UpdateUserVlanid2(user.Username, reqvid2)
	}
}
