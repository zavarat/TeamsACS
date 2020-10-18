package radiusd

import (
	"fmt"
)

func (s *AuthService) CheckOnlineCount(username string, activeNum int) error {
	if activeNum != 0 {
		onlineCount, _ := s.Manager.GetRadiusManager().GetOnlineCount(username)
		if int(onlineCount) > activeNum {
			return fmt.Errorf("user:%s active num over limit(max=%d)", username, activeNum)
		}
	}
	return nil
}
