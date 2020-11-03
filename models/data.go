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

	"github.com/ca17/teamsacs/common/web"
)

// A generic data CRUD management API with no predefined schema,
// storing extra data that you may not use at all, but you'll probably use a lot.

// DataManager
type DataManager struct{ *ModelManager }

func (m *ModelManager) GetDataManager() *DataManager {
	store, _ := m.ManagerMap.Get("DataManager")
	return store.(*DataManager)
}

// QueryDatas
func (m *DataManager) QueryDatas(params web.RequestParams) (*web.PageResult, error) {
	collname := params.GetMustString("collname")
	return m.QueryPagerItems(params, collname)
}


// GetDataById
func (m *DataManager) GetData(params web.RequestParams) (*Attributes, error) {
	_id := params.GetMustString("_id")
	collname := params.GetMustString("collname")
	coll := m.GetTeamsAcsCollection(collname)
	doc := coll.FindOne(context.TODO(), bson.M{"_id": _id})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Attributes)
	err = doc.Decode(result)
	return result, err
}


// AddData
func (m *DataManager) AddData(params web.RequestParams) error {
	data := params.GetParamMap("data")
	_ = data.GetMustString("_id")
	coll := m.GetTeamsAcsCollection(params.GetMustString("collname"))
	_, err := coll.InsertOne(context.TODO(), data)
	return err
}

// UpdateData
func (m *DataManager) UpdateData(params web.RequestParams) error {
	_id := params.GetMustString("_id")
	query := bson.M{"_id": _id}
	update := bson.M{"$set": params.GetParamMap("data")}
	_, err := m.GetTeamsAcsCollection(params.GetMustString("collname")).UpdateOne(context.TODO(), query, update)
	return err
}

// DeleteData
func (m *DataManager) DeleteData(params web.RequestParams) error {
	_id := params.GetMustString("_id")
	collname := params.GetMustString("collname")
	_, err := m.GetTeamsAcsCollection(collname).DeleteOne(context.TODO(), bson.M{"_id": _id})
	return err
}

