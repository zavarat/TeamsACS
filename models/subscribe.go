/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubscribeManager struct{ *ModelManager }

func (m *ModelManager) GetSubscribeManager() *SubscribeManager {
	store, _ := m.ManagerMap.Get("SubscribeManager")
	return store.(*SubscribeManager)
}

func (m *SubscribeManager) GetSubscribeByUser(username string) (*Subscribe, error) {
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


func (m *SubscribeManager) GetSubscribeByMac(mac string) (*Subscribe, error) {
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


func (m *SubscribeManager) UpdateSubscribeByUsername(username string, valmap map[string]interface{})  error {
	coll := m.GetTeamsAcsCollection(TeamsacsSubscribe)
	_, err := coll.UpdateOne(context.TODO(), bson.M{"username":username}, valmap)
	return err
}

func (m *SubscribeManager) AddSubscribe(subs *Subscribe) (string, error) {
	r, err := m.GetTeamsAcsCollection(TeamsacsSubscribe).InsertOne(context.TODO(), subs)
	if err != nil {
		return "", err
	}
	return r.InsertedID.(primitive.ObjectID).Hex(), err
}
