package radiusd

import (
	"context"
	"fmt"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"layeh.com/radius/rfc2866"

	"github.com/ca17/teamsacs/models"
)

func (s *AcctService) DoAcctNasOn(r *radius.Request) {
	err := s.Manager.GetRadiusManager().BatchClearRadiusOnlineDataByNas(
		rfc2865.NASIPAddress_Get(r.Packet).String(),
		rfc2865.NASIdentifier_GetString(r.Packet),
	)
	if err != nil {
		radlog.Errorf("BatchClearRadiusOnlineDataByNas error, %s", err.Error())
	}
}

func (s *AcctService) DoAcctNasOff(r *radius.Request) {
	err := s.Manager.GetRadiusManager().BatchClearRadiusOnlineDataByNas(
		rfc2865.NASIPAddress_Get(r.Packet).String(),
		rfc2865.NASIdentifier_GetString(r.Packet),
	)
	if err != nil {
		radlog.Errorf("BatchClearRadiusOnlineDataByNas error, %s", err.Error())
	}
}

func (s *AcctService) DoAcctDisconnect(r *radius.Request, vpe *models.Vpe, username, nasrip string) {
	packet := radius.New(radius.CodeDisconnectRequest, []byte(vpe.Secret))
	sessionid := rfc2866.AcctSessionID_GetString(r.Packet)
	if sessionid == "" {
		radlog.Errorf("radius disconnect user:%s, but sessionid is empty", username)
		return
	}
	_ = rfc2865.UserName_SetString(packet, username)
	_ = rfc2866.AcctSessionID_Set(packet, []byte(sessionid))
	radlog.Infof("disconnect user:%s => (%s:%d): %s", username, nasrip, vpe.CoaPort, FormatPacket(packet))
	response, err := radius.Exchange(context.Background(), packet, fmt.Sprintf("%s:%d", nasrip, vpe.CoaPort))
	if err != nil {
		radlog.Errorf("radius disconnect user:%s failure", username)
	}
	radlog.Info("radius disconnect resp from (%s:%s): %s ", nasrip, vpe.CoaPort, FormatPacket(response))
}
