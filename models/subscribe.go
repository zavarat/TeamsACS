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


// Subscribe
type Subscribe struct {
	ID         string      `bson:"_id,omitempty" json:"id,omitempty"`
	VpeSids    []string    `bson:"vpe_sids,omitempty" json:"vpe_sids,omitempty"`
	Profile    ProfileAttr `bson:"profile,omitempty" json:"profile,omitempty,omitempty"`
	Realname   string      `bson:"realname,omitempty" json:"realname,omitempty"`
	Email      string      `bson:"email,omitempty" json:"email,omitempty"`
	Username   string      `bson:"username,omitempty" json:"username,omitempty"`
	Password   string      `bson:"password,omitempty" json:"password,omitempty"`
	Ipaddr     string      `bson:"ipaddr,omitempty" json:"ipaddr,omitempty"`
	Macaddr    string      `bson:"macaddr,omitempty" json:"macaddr,omitempty"`
	Vlanid1    int         `bson:"vlanid_1,omitempty" json:"vlanid1,omitempty"`
	Vlanid2    int         `bson:"vlanid_2,omitempty" json:"vlanid2,omitempty"`
	BindMac    int         `bson:"bind_mac,omitempty" json:"bind_mac,omitempty"`
	BindVlan   int         `bson:"bind_vlan,omitempty" json:"bind_vlan,omitempty"`
	Status     string      `bson:"status,omitempty" json:"status,omitempty"`
	Remark     string      `bson:"remark,omitempty" json:"remark,omitempty"`
	ExpireTime time.Time   `bson:"expire_time,omitempty" json:"expire_time,omitempty"`
	Timestamp  time.Time   `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
}


type SubscribeManager struct{ *ModelManager }

func (m *ModelManager) GetSubscribeManager() *SubscribeManager {
	store, _ := m.ManagerMap.Get("SubscribeManager")
	return store.(*SubscribeManager)
}

// QuerySubscribes
func (m *SubscribeManager) QuerySubscribes(params web.RequestParams) (*web.PageResult, error) {
	return m.QueryPagerItems(params, TeamsacsSubscribe)
}

// GetSubscribeByUser
func (m *SubscribeManager) GetSubscribeByUser(username string) (*Subscribe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsSubscribe)
	doc := coll.FindOne(context.TODO(), bson.M{"username": username})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Subscribe)
	err = doc.Decode(result)
	return result, err
}

// GetSubscribeByMac
func (m *SubscribeManager) GetSubscribeByMac(mac string) (*Subscribe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsSubscribe)
	doc := coll.FindOne(context.TODO(), bson.M{"macaddr": mac})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Subscribe)
	err = doc.Decode(result)
	return result, err
}

// UpdateSubscribeByUsername
func (m *SubscribeManager) UpdateSubscribeByUsername(username string, valmap map[string]interface{}) error {
	coll := m.GetTeamsAcsCollection(TeamsacsSubscribe)
	_, err := coll.UpdateOne(context.TODO(), bson.M{"username": username}, valmap)
	return err
}

// ExistSubscribe
func (m *SubscribeManager) ExistSubscribe(username string) bool {
	coll := m.GetTeamsAcsCollection(TeamsacsSubscribe)
	count, _ := coll.CountDocuments(context.TODO(), bson.M{"username": username})
	return count > 0
}

// UpdateSubscribe
// update by username
func (m *SubscribeManager) UpdateSubscribe(subscribe *Subscribe) error {
	if err := subscribe.UpdateValidate(); err != nil {
		return err
	}
	coll := m.GetTeamsAcsCollection(TeamsacsSubscribe)
	query := bson.M{"username": subscribe.Username}
	update := bson.M{"$set": subscribe}
	_, err := coll.UpdateOne(context.TODO(), query, update)
	return err
}

// AddSubscribe
func (m *SubscribeManager) AddSubscribe(subs *Subscribe) (string, error) {
	if err := subs.AddValidate(); err != nil {
		return "", err
	}
	if m.ExistSubscribe(subs.Username) {
		return "", fmt.Errorf("subscribe exists")
	}
	subs.Status = constant.ENABLED
	r, err := m.GetTeamsAcsCollection(TeamsacsSubscribe).InsertOne(context.TODO(), subs)
	if err != nil {
		return "", err
	}
	return r.InsertedID.(string), err
}

// DeleteSubscribe
func (m *SubscribeManager) DeleteSubscribe(username string) error {
	if common.IsEmptyOrNA(username) {
		return fmt.Errorf("username is empty or NA")
	}
	_, err := m.GetTeamsAcsCollection(TeamsacsSubscribe).DeleteOne(context.TODO(), bson.M{"username": username})
	return err
}
