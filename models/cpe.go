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

// Cpe
type Cpe = DataObject

type CpeManager struct{ *ModelManager }

func (m *ModelManager) GetCpeManager() *CpeManager {
	store, _ := m.ManagerMap.Get("CpeManager")
	return store.(*CpeManager)
}

func (m *CpeManager) GetCpeByDeviceId(device_id string) (*Cpe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsCpe)
	doc := coll.FindOne(context.TODO(), bson.M{"device_id": device_id})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Cpe)
	err = doc.Decode(result)
	return result, err
}

func (m *CpeManager) GetCpeBySn(sn string) (*Cpe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsCpe)
	doc := coll.FindOne(context.TODO(), bson.M{"sn": sn})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Cpe)
	err = doc.Decode(result)
	return result, err
}

func (m *CpeManager) QueryCpes(params web.RequestParams) (*web.PageResult, error) {
	return m.QueryPagerItems(params, TeamsacsCpe)
}

func (m *CpeManager) ExistCpe(sn string) bool {
	coll := m.GetTeamsAcsCollection(TeamsacsCpe)
	count, _ := coll.CountDocuments(context.TODO(), bson.M{"sn": sn})
	return count > 0
}
