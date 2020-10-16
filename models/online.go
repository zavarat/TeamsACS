package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)



func (m *RadiusManager) GetOnlineCount(username string) (int64, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsOnline)
	return coll.CountDocuments(context.TODO(), bson.M{"username": username})
}

