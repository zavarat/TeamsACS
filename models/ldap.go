package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LdapManager struct{ *ModelManager }

func (m *ModelManager) GetLdapManager() *LdapManager {
	store, _ := m.ManagerMap.Get("LdapManager")
	return store.(*LdapManager)
}

func (m *LdapManager) FindLdapById(id primitive.ObjectID) (*Ldap, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsLdap)
	doc := coll.FindOne(context.TODO(), bson.M{"_id":id})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Ldap)
	err = doc.Decode(result)
	return result, err
}


