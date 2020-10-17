package radiusd

import (
	"errors"

	"github.com/ca17/teamsacs/models"
	"github.com/ca17/teamsacs/radiusd/radparser"
)

// mac binding detection
// Detected only if both user mac and request mac are valid.
// If user mac is empty, update user mac directly.
func (s *AuthService) CheckMacBind(user *models.Subscribe, vendorReq *radparser.VendorRequest) error {
	if user.BindMac == 1 {
		if user.Macaddr != "" && user.Macaddr != "N/A" && vendorReq.Macaddr != "" {
			if user.Macaddr != vendorReq.Macaddr {
				return errors.New("user mac bind not match")
			}
		}
	}
	return nil
}
