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
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/ca17/teamsacs/common/web"
	"github.com/ca17/teamsacs/constant"
)

type Subscribe = DataObject

// AuthorizationProfile Method
func (a Subscribe) GetExpireTime() time.Time {
	return a.GetDateValue("expire_time", time.Now().Add(time.Second*60))
}

func (a Subscribe) GetInterimInterval() int {
	return a.GetIntValue("interim_interval", 120)
}

func (a Subscribe) GetAddrPool() string {
	return a.GetStringValue("addr_pool", constant.NA)
}

func (a Subscribe) GetIpaddr() string {
	return a.GetStringValue("ipaddr", constant.NA)
}

func (a Subscribe) GetUpRateKbps() int {
	return a.GetIntValue("up_rate", 0)
}

func (a Subscribe) GetDownRateKbps() int {
	return a.GetIntValue("down_rate", 0)
}

func (a Subscribe) GetDomain() string {
	return a.GetStringValue("domain", constant.NA)
}

func (a Subscribe) GetLimitPolicy() string {
	return a.GetStringValue("limit_policy", constant.NA)
}

func (a Subscribe) GetUpLimitPolicy() string {
	return a.GetStringValue("up_limit_policy", constant.NA)
}

func (a Subscribe) GetDownLimitPolicy() string {
	return a.GetStringValue("down_limit_policy", constant.NA)
}


func (a Subscribe) GetMacAddr() string {
	return a.GetStringValue("mac_addr", constant.NA)
}

func (a Subscribe) GetPassword() string {
	return a.GetStringValue("password", constant.NA)
}

func (a Subscribe) GetUsername() string {
	return a.GetStringValue("username", constant.NA)
}

func (a Subscribe) GetActiveNum() int {
	return a.GetIntValue("active_num", 0)
}

func (a Subscribe) GetStatus() string {
	return a.GetStringValue("status", constant.DISABLED)
}




// SubscribeManager
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
