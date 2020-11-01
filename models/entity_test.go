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
	"encoding/json"
	"testing"
	"time"

	"github.com/ca17/teamsacs/common"
)

func TestOperator2json(t *testing.T) {
	item := &Operator{
		ID:        common.UUID(),
		Email:     "test@teamsacs.com",
		Username:  "opr",
		Level:     "opr",
		Remark:    "opr",
	}
	bs, err := json.MarshalIndent(item, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bs))

}

func TestCpe2json(t *testing.T) {
	cpe := &Cpe{
		Id:         common.UUID(),
		Sn:         "xxxxxx",
		DeviceId:   "",
		Attrs:      nil,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	bs, err := json.MarshalIndent(cpe, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bs))

}

func TestSubscribe2json(t *testing.T) {
	item := &Subscribe{
		ID:         common.UUID(),
		VpeSids:    nil,
		Profile:    ProfileAttr{
			Domain:          "N/A",
			InterimInterval: 0,
			AddrPool:        "N/A",
			ActiveNum:       1,
			MfaStatus:       "disabled",
			UpRate:          1048576,
			DownRate:        1048576,
			LimitPolicy:     "N/A",
			UpLimitPolicy:   "N/A",
			DownLimitPolicy: "N/A",
		},
		Realname:   "test",
		Email:      "test@teamsacs.com",
		Username:   "test01",
		Password:   "888888",
		MfaStatus:  "disabled",
		MfaSecret:  "N/A",
		Ipaddr:     "N/A",
		Macaddr:    "N/A",
		Vlanid1:    0,
		Vlanid2:    0,
		BindMac:    0,
		BindVlan:   0,
		Status:     "Enabled",
		Remark:     "Test user",
		ExpireTime: time.Now(),
		Timestamp:  time.Now(),
	}
	bs, err := json.MarshalIndent(item,"","\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bs))

}


