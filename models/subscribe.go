package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type SubscribeManager struct{ *ModelManager }

func (m *ModelManager) GetSubscribeManager() *SubscribeManager {
	store, _ := m.ManagerMap.Get("SubscribeManager")
	return store.(*SubscribeManager)
}

func (m *SubscribeManager) FindSubscribeByUser(username string) (*Subscribe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsSubscribe)
	doc := coll.FindOne(context.TODO(), bson.M{"username":username})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Subscribe)
	err = doc.Decode(result)
	return result, err
}


func (m *SubscribeManager) FindSubscribeByMac(mac string) (*Subscribe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsSubscribe)
	doc := coll.FindOne(context.TODO(), bson.M{"macaddr":mac})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Subscribe)
	err = doc.Decode(result)
	return result, err
}


func (m *SubscribeManager) UpdateSubscribeByUser(username string, valmap map[string]interface{})  error {
	coll := m.GetTeamsAcsCollection(TeamsacsSubscribe)
	_, err := coll.UpdateOne(context.TODO(), bson.M{"username":username}, valmap)
	return err
}
