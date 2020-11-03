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

	"go.mongodb.org/mongo-driver/bson"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/web"
	"github.com/ca17/teamsacs/constant"
)

// Vpe
// VPE is also a BRAS system
type Vpe struct {
	ID         string     `bson:"_id,omitempty" json:"id,omitempty"`
	Sn         string     `bson:"sn" json:"sn,omitempty"`
	DeviceId   string     `bson:"device_id" json:"device_id,omitempty" `
	Attrs      Attributes `bson:"attrs" json:"attrs,omitempty" `
	Identifier string     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name       string     `bson:"name,omitempty" json:"name,omitempty"`
	Ipaddr     string     `bson:"ipaddr,omitempty" json:"ipaddr,omitempty"`
	Secret     string     `bson:"secret,omitempty" json:"secret,omitempty"`
	VendorCode string     `bson:"vendor_code,omitempty" json:"vendor_code,omitempty"`
	CoaPort    int        `bson:"coa_port,omitempty" json:"coa_port,omitempty"`
	Status     string     `bson:"status,omitempty" json:"status,omitempty"`
	Remark     string     `bson:"remark,omitempty" json:"remark,omitempty"`
}



func (a *Vpe) AddValidate() error {
	switch {
	case common.IsEmptyOrNA(a.Sn):
		return fmt.Errorf("invalid sn")
	case common.IsEmptyOrNA(a.Identifier):
		return fmt.Errorf("invalid identifier")
	case common.IsEmptyOrNA(a.Ipaddr):
		return fmt.Errorf("invalid ipaddr")
	case common.IsEmptyOrNA(a.Secret):
		return fmt.Errorf("invalid secret")
	case common.IsEmptyOrNA(a.VendorCode):
		return fmt.Errorf("invalid vendor_code")
	}
	return nil
}


// VpeManager
type VpeManager struct{ *ModelManager }

func (m *ModelManager) GetVpeManager() *VpeManager {
	store, _ := m.ManagerMap.Get("VpeManager")
	return store.(*VpeManager)
}

// QueryVpes
func (m *VpeManager) QueryVpes(params web.RequestParams) (*web.PageResult, error) {
	return m.QueryPagerItems(params, TeamsacsVpe)
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
func (m *VpeManager) UpdateVpe(vpe *Vpe) error {
	if common.IsEmptyOrNA(vpe.Sn) {
		return fmt.Errorf("sn is empty or NA")
	}

	query := bson.M{"sn": vpe.Sn}
	update := bson.M{"$set": vpe}
	_, err := m.GetTeamsAcsCollection(TeamsacsVpe).UpdateOne(context.TODO(), query, update)
	return err
}

// DeleteVpe
func (m *VpeManager) DeleteVpe(sn string) error {
	if common.IsEmptyOrNA(sn) {
		return fmt.Errorf("sn is empty or NA")
	}
	_, err := m.GetTeamsAcsCollection(TeamsacsVpe).DeleteOne(context.TODO(), bson.M{"sn": sn})
	return err
}
