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
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/web"
	"github.com/ca17/teamsacs/constant"
)

type VpeManager struct{ *ModelManager }

func (m *ModelManager) GetVpeManager() *VpeManager {
	store, _ := m.ManagerMap.Get("VpeManager")
	return store.(*VpeManager)
}

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

func (m *VpeManager) AddVpe(form *web.WebForm) error {
	coll := m.GetTeamsAcsCollection(TeamsacsVpe)
	vpe := &Vpe{
		Identifier: form.GetVal2("identifier", constant.NA),
		Sn:         form.GetVal2("sn", constant.NA),
		Name:       form.GetVal2("name", constant.NA),
		Ipaddr:     form.GetVal2("ipaddr", constant.NA),
		Secret:     form.GetVal2("secret", constant.NA),
		VendorCode: form.GetVal2("vendor_code", constant.NA),
		CoaPort:    form.GetIntVal("coa_port", 3799),
		Status:     form.GetVal2("status", constant.ENABLED),
		Remark:     form.GetVal2("remark", constant.NA),
	}
	vpe.Attrs = bsonx.Doc{}
	for k, v := range form.Posts {
		 if !strings.HasPrefix(k, "attrs.") {
			continue
		}
		vpe.Attrs.Set(k[6:], bsonx.String(v[0]))
	}
	_, err := coll.InsertOne(context.TODO(), vpe)
	return err
}

// UpdateCpe
// Update VPE information. Undefined properties are not accepted, but the attrs property can be modified at will.
func (m *VpeManager) UpdateCpe(form *web.WebForm) error {
	sn := form.GetVal("sn")
	if common.IsEmptyOrNA(sn) {
		return fmt.Errorf("sn is empty or NA")
	}

	data := bson.M{
		"update_time": primitive.NewDateTimeFromTime(time.Now()),
	}
	for k, v := range form.Posts {
		if k == "sn" || (!common.InSlice(k,[]string{
			"identifier","name","ipaddr","secret","vendor_code","coa_port","remark","status", "ldap_id"}) &&
				!strings.HasPrefix(k,"attrs.")){
			continue
		}
		data[k] = v[0]
	}
	query := bson.M{"sn": sn}
	update := bson.M{"$set": data}
	_, err := m.GetTeamsAcsCollection(TeamsacsVpe).UpdateOne(context.TODO(), query, update)
	return err
}


func (m *CpeManager) DeleteVpe(sn string) error {
	if common.IsEmptyOrNA(sn) {
		return fmt.Errorf("sn is empty or NA")
	}
	_, err := m.GetTeamsAcsCollection(TeamsacsVpe).DeleteOne(context.TODO(), bson.M{"sn": sn})
	return err
}
