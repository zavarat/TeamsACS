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
	"time"
)

type Attributes = map[string]interface{}


// VariableConfig
type VariableConfig struct {
	ID     string     `bson:"_id,omitempty" json:"id,omitempty"`
	Vendor string     `bson:"vendor" json:"vendor,omitempty"`
	Group  string     `bson:"group" json:"group,omitempty"`
	Attrs  Attributes `bson:"attrs" json:"attrs,omitempty"`
	Remark string     `bson:"remark" json:"remark,omitempty"`
}

// AppTemplate
// Application template, defining an application specification
type AppTemplate struct {
	ID    string     `bson:"_id,omitempty" json:"id,omitempty"`
	Attrs Attributes `bson:"attrs" json:"attrs,omitempty"`
}

// AppGroup
// Creating a set of application specifications through application templates
type AppGroup struct {
	ID   string       `bson:"_id,omitempty" json:"id,omitempty"`
	Name string       `bson:"name,omitempty" json:"name,omitempty"`
	Apps []Attributes `bson:"apps" json:"apps,omitempty"`
}



// ProfileAttr
// Radius profile attrs
type ProfileAttr struct {
	Domain          string `bson:"domain" json:"domain,omitempty"`
	InterimInterval int    `bson:"interim_interval" json:"interim_interval,omitempty"`
	AddrPool        string `bson:"addr_pool" json:"addr_pool,omitempty"`
	ActiveNum       int    `bson:"active_num" json:"active_num,omitempty"`
	UpRate          int    `bson:"up_rate" json:"up_rate,omitempty"`
	DownRate        int    `bson:"down_rate" json:"down_rate,omitempty"`
	LimitPolicy     string `bson:"limit_policy" json:"limit_policy,omitempty"`
	UpLimitPolicy   string `bson:"up_limit_policy" json:"up_limit_policy,omitempty"`
	DownLimitPolicy string `bson:"down_limit_policy" json:"down_limit_policy,omitempty"`
}

// Profile
// Radius profile
type Profile struct {
	ProfileAttr
	ID           string `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string `bson:"name" json:"name,omitempty"`
	BillTimes    int    `bson:"bill_times" json:"bill_times,omitempty"`
	BillTimeunit string `bson:"bill_timeunit" json:"bill_timeunit,omitempty"`
	Status       string `bson:"status" json:"status,omitempty"`
	Remark       string `bson:"remark" json:"remark,omitempty"`
}


// AuthorizationProfile Method

func (a Subscribe) GetExpireTime() time.Time {
	return a.ExpireTime
}

func (a Subscribe) GetInterimInterval() int {
	return a.Profile.InterimInterval
}

func (a Subscribe) GetAddrPool() string {
	return a.Profile.AddrPool
}

func (a Subscribe) GetIpaddr() string {
	return a.Ipaddr
}

func (a Subscribe) GetUpRateKbps() int {
	return a.Profile.UpRate
}

func (a Subscribe) GetDownRateKbps() int {
	return a.Profile.DownRate
}

func (a Subscribe) GetDomain() string {
	return a.Profile.Domain
}

func (a Subscribe) GetLimitPolicy() string {
	return a.Profile.LimitPolicy
}

func (a Subscribe) GetUpLimitPolicy() string {
	return a.Profile.UpLimitPolicy
}

func (a Subscribe) GetDownLimitPolicy() string {
	return a.Profile.DownLimitPolicy
}
