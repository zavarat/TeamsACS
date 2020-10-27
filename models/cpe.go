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
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/web"
)

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

func (m *CpeManager) QueryCpes(form *web.WebForm) (*web.PageResult, error) {
	var findOptions = options.Find()
	var pos = form.GetInt64Val("start", 0)
	findOptions.SetSkip(pos)
	findOptions.SetLimit(form.GetInt64Val("count", 40))
	coll := m.GetTeamsAcsCollection(TeamsacsCpe)
	var q = bson.M{}
	for qname, vals := range form.Gets {
		if !strings.HasPrefix(qname, "attr_") {
			continue
		}
		q["attrs."+qname[5:]] = vals[0]
	}
	cur, err := coll.Find(context.TODO(), q, findOptions)
	if err != nil {
		return nil, err
	}
	var countOptions = options.Count()
	total, err := coll.CountDocuments(context.TODO(), q, countOptions)
	if err != nil {
		return nil, err
	}
	items := make([]map[string]interface{}, 0)
	for cur.Next(context.TODO()) {
		var elem map[string]interface{}
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		} else {
			items = append(items, elem)
		}
	}
	return &web.PageResult{TotalCount: total, Pos: pos, Data: items}, nil
}

func (m *CpeManager) AddCpe(form *web.WebForm) error {
	sn := form.GetVal("sn")
	if common.IsEmptyOrNA(sn) {
		return fmt.Errorf("sn is empty or NA")
	}
	coll := m.GetTeamsAcsCollection(TeamsacsCpe)
	cpe := new(Cpe)
	cpe.Sn = sn
	cpe.DeviceId = ""
	cpe.CreateTime = primitive.NewDateTimeFromTime(time.Now())
	cpe.UpdateTime = primitive.NewDateTimeFromTime(time.Now())
	cpe.Attrs = bsonx.Doc{}
	for k, v := range form.Posts {
		cpe.Attrs.Set(k, bsonx.String(v[0]))
	}
	_, err := coll.InsertOne(context.TODO(), cpe)
	return err
}

func (m *CpeManager) UpdateCpeAttrs(form *web.WebForm) error {
	sn := form.GetVal("sn")
	if common.IsEmptyOrNA(sn) {
		return fmt.Errorf("sn is empty or NA")
	}
	data := bson.M{
		"update_time": primitive.NewDateTimeFromTime(time.Now()),
	}
	for k, v := range form.Posts {
		data["attrs."+k] = v[0]
	}
	query := bson.M{"sn": sn}
	update := bson.M{"$set": data}
	_, err := m.GetTeamsAcsCollection(TeamsacsCpe).UpdateOne(context.TODO(), query, update)
	return err
}

func (m *CpeManager) DeleteCpe(sn string) error {
	if common.IsEmptyOrNA(sn) {
		return fmt.Errorf("sn is empty or NA")
	}
	_, err := m.GetTeamsAcsCollection(TeamsacsCpe).DeleteOne(context.TODO(), bson.M{"sn": sn})
	return err
}
