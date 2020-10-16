package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type VpeManager struct{ *ModelManager }

func (m *ModelManager) GetVpeManager() *VpeManager {
	store, _ := m.ManagerMap.Get("VpeManager")
	return store.(*VpeManager)
}

func (m *VpeManager) FindVpeByIpaddr(ip string) (*Vpe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsVpe)
	doc := coll.FindOne(context.TODO(), bson.M{"ipaddr":ip})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Vpe)
	err = doc.Decode(result)
	return result, err
}


func (m *VpeManager) FindVpeByIdentifier(identifier string) (*Vpe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsVpe)
	doc := coll.FindOne(context.TODO(), bson.M{"identifier":identifier})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Vpe)
	err = doc.Decode(result)
	return result, err
}


