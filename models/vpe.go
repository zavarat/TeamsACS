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
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/web"
	"github.com/ca17/teamsacs/constant"
)

type VpeManager struct{ *ModelManager }

func (m *ModelManager) GetVpeManager() *VpeManager {
	store, _ := m.ManagerMap.Get("VpeManager")
	return store.(*VpeManager)
}

// GetVpeByIpaddr
func (m *VpeManager) GetVpeByIpaddr(ip string) (*Vpe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsVpe)
	doc := coll.FindOne(context.TODO(), bson.M{"ipaddr": ip})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Vpe)
	err = doc.Decode(result)
	return result, err
}

// GetVpeByIdentifier
func (m *VpeManager) GetVpeByIdentifier(identifier string) (*Vpe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsVpe)
	doc := coll.FindOne(context.TODO(), bson.M{"identifier": identifier})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Vpe)
	err = doc.Decode(result)
	return result, err
}

// ExistVpe
func (m *VpeManager) ExistVpe(sn string) bool {
	coll := m.GetTeamsAcsCollection(TeamsacsVpe)
	count, _ := coll.CountDocuments(context.TODO(), bson.M{"sn": sn})
	return count > 0
}


// AddVpe
func (m *VpeManager) AddVpe(item *Vpe) error {
	if err := item.AddValidate(); err!=nil {
		return err
	}
	item.Status = constant.ENABLED
	coll := m.GetTeamsAcsCollection(TeamsacsVpe)
	if m.ExistVpe(item.Sn) {
		return fmt.Errorf("vpe exists")
	}
	_, err := coll.InsertOne(context.TODO(), item)
	return err
}

// UpdateVpe
// Update VPE information. Undefined properties are not accepted, but the attrs property can be modified at will.
func (m *VpeManager) UpdateVpe(params web.RequestParams) error {
	sn := params.GetString("sn")
	if common.IsEmptyOrNA(sn) {
		return fmt.Errorf("sn is empty or NA")
	}
	data := bson.M{
		"update_time": time.Now(),
	}
	for k, v := range params.GetParamMap("attrs") {
		data["attrs."+k] = v
	}
	query := bson.M{"sn": sn}
	update := bson.M{"$set": data}
	_, err := m.GetTeamsAcsCollection(TeamsacsVpe).UpdateOne(context.TODO(), query, update)
	return err
}

// DeleteVpe
func (m *CpeManager) DeleteVpe(sn string) error {
	if common.IsEmptyOrNA(sn) {
		return fmt.Errorf("sn is empty or NA")
	}
	_, err := m.GetTeamsAcsCollection(TeamsacsVpe).DeleteOne(context.TODO(), bson.M{"sn": sn})
	return err
}
