package radiusd

import (
	"fmt"
)

func (s *AuthService) CheckOnlineCount(username string, activeNUm int) error {
	if activeNUm != 0 {
		onlineCount, _ := s.Manager.GetRadiusManager().GetOnlineCount(username)
		if int(onlineCount) > activeNUm {
			return fmt.Errorf("user:%s active num over limit(max=%d)", username, activeNUm)
		}
	}
	return nil
}
