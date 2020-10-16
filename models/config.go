package models

import (
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

type ConfigManager struct{ *ModelManager }

func (m *ModelManager) GetConfigManager() *ConfigManager {
	store, _ := m.ManagerMap.Get("ConfigManager")
	return store.(*ConfigManager)
}


func (m *ConfigManager) GetConfigValue(ctype, name string) string{
	coll := m.GetTeamsAcsCollection(TeamsacsVpe)
	doc := coll.FindOne(context.TODO(), bson.M{"type":ctype, "name":name})
	err := doc.Err()
	if err != nil {
		return ""
	}
	var result = new(Config)
	err = doc.Decode(&result)
	return result.Value
}

func (m *ConfigManager) GetRadiusConfigValue(name string) string{
	coll := m.GetTeamsAcsCollection(TeamsacsVpe)
	doc := coll.FindOne(context.TODO(), bson.M{"type":"radius", "name":name})
	err := doc.Err()
	if err != nil {
		return ""
	}
	var result = ""
	err = doc.Decode(&result)
	return result
}

func (m *ConfigManager) GetRadiusConfigStringValue(name string, defval string) string{
	val := m.GetRadiusConfigValue(name)
	if val == "" {
		return defval
	}
	return val
}

func (m *ConfigManager) GetRadiusConfigIntValue(name string, defval int64) int64{
	val := m.GetRadiusConfigValue(name)
	if val == "" {
		return defval
	}
	v, err :=  strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defval
	}
	return v
}
