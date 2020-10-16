package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RadiusManager struct{ *ModelManager }

func (m *ModelManager) GetRadiusManager() *RadiusManager {
	store, _ := m.ManagerMap.Get("RadiusManager")
	return store.(*RadiusManager)
}

func (m *RadiusManager) AddRadiusAuthLog(username string, nasip string, result string, reason string, cast int64) error {
	authlog := Authlog{
		Username:  username,
		NasAddr:   nasip,
		Result:    result,
		Reason:    reason,
		Cast:      int(cast),
		Timestamp: primitive.NewDateTimeFromTime(time.Now()),
	}
	coll := m.GetTeamsAcsCollection(TeamsacsAuthlog)
	_, err := coll.InsertOne(context.TODO(), authlog)
	return err
}

func (m *RadiusManager) BatchClearRadiusOnlineDataByNas(nasip, nasid string) error {
	coll := m.GetTeamsAcsCollection(TeamsacsOnline)
	filter := bson.D{
		{"$or",
			bson.A{
				bson.D{{"nas_addr", nasip}},
				bson.D{{"nas_id", nasid}},
			}},
	}
	_, err := coll.DeleteMany(context.TODO(), filter)
	return err
}

func (m *RadiusManager) AddRadiusOnline(ol Accounting) error {
	_, err := m.GetTeamsAcsCollection(TeamsacsOnline).InsertOne(context.TODO(), ol)
	return err
}

func (m *RadiusManager) AddRadiusAccounting(acct Accounting) error {
	acct.AcctStopTime = primitive.NewDateTimeFromTime(time.Now())
	_, err := m.GetTeamsAcsCollection(TeamsacsAccounting).InsertOne(context.TODO(), acct)
	return err
}

func (m *RadiusManager) DeleteRadiusOnline(sessionid string) error {
	_, err := m.GetTeamsAcsCollection(TeamsacsOnline).DeleteOne(context.TODO(), bson.M{"acct_session_id": sessionid})
	return err
}


func (m *RadiusManager) UpdateRadiusOnline(acct Accounting) error {
	data := bson.D{
		{"$inc", bson.D{
			{"acct_input_total", acct.AcctInputTotal},
			{"acct_output_total", acct.AcctOutputTotal},
			{"acct_input_packets", acct.AcctInputPackets},
			{"acct_output_packets", acct.AcctOutputPackets},
			{"acct_input_total", acct.AcctSessionTime},
		}},
		{"last_update", primitive.NewDateTimeFromTime(time.Now())},
	}
	query := bson.M{"acct_session_id": acct.AcctSessionId}
	r := m.GetTeamsAcsCollection(TeamsacsOnline).FindOne(context.TODO(), query)
	if r.Err() == nil {
		return m.AddRadiusAccounting(acct)
	}
	_, err := m.GetTeamsAcsCollection(TeamsacsOnline).UpdateOne(context.TODO(), query, data)
	return err
}
