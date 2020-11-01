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
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/log"
	"github.com/ca17/teamsacs/common/web"
)

type RadiusManager struct{ *ModelManager }

func (m *ModelManager) GetRadiusManager() *RadiusManager {
	store, _ := m.ManagerMap.Get("RadiusManager")
	return store.(*RadiusManager)
}

func (m *RadiusManager) AddRadiusAuthLog(username string, nasip string, result string, reason string, cast int64) error {
	authlog := Authlog{
		Username:  username,
		NasAddr:   nasip,
		Result:    result,
		Reason:    reason,
		Cast:      int(cast),
		Timestamp: time.Now(),
	}
	coll := m.GetTeamsAcsCollection(TeamsacsAuthlog)
	_, err := coll.InsertOne(context.TODO(), authlog)
	return err
}

func (m *RadiusManager) BatchClearRadiusOnlineDataByNas(nasip, nasid string) error {
	coll := m.GetTeamsAcsCollection(TeamsacsOnline)
	filter := bson.D{
		{"$or",
			bson.A{
				bson.D{{"nas_addr", nasip}},
				bson.D{{"nas_id", nasid}},
			}},
	}
	_, err := coll.DeleteMany(context.TODO(), filter)
	return err
}

func (m *RadiusManager) AddRadiusOnline(ol Accounting) error {
	_, err := m.GetTeamsAcsCollection(TeamsacsOnline).InsertOne(context.TODO(), ol)
	return err
}

func (m *RadiusManager) AddRadiusAccounting(acct Accounting) error {
	acct.AcctStopTime = time.Now()
	_, err := m.GetTeamsAcsCollection(TeamsacsAccounting).InsertOne(context.TODO(), acct)
	return err
}

func (m *RadiusManager) DeleteRadiusOnline(sessionid string) error {
	_, err := m.GetTeamsAcsCollection(TeamsacsOnline).DeleteOne(context.TODO(), bson.M{"acct_session_id": sessionid})
	return err
}


func (m *RadiusManager) UpdateRadiusOnlineData(acct Accounting) error {
	data := bson.D{
		{"$inc", bson.D{
			{"acct_input_total", acct.AcctInputTotal},
			{"acct_output_total", acct.AcctOutputTotal},
			{"acct_input_packets", acct.AcctInputPackets},
			{"acct_output_packets", acct.AcctOutputPackets},
			{"acct_input_total", acct.AcctSessionTime},
		}},
		{"last_update", primitive.NewDateTimeFromTime(time.Now())},
	}
	query := bson.M{"acct_session_id": acct.AcctSessionId}
	r := m.GetTeamsAcsCollection(TeamsacsOnline).FindOne(context.TODO(), query)
	if r.Err() == nil {
		return m.AddRadiusAccounting(acct)
	}
	_, err := m.GetTeamsAcsCollection(TeamsacsOnline).UpdateOne(context.TODO(), query, data)
	return err
}


func getAcctStartTime(sessionTime string) time.Time {
	m, _ := time.ParseDuration("-" + sessionTime + "s")
	return time.Now().Add(m)
}

func getInputTotal(form *web.WebForm) int64 {
	var acctInputOctets = form.GetInt64Val("acctInputOctets", 0)
	var acctInputGigawords = form.GetInt64Val("acctInputGigaword", 0)
	return acctInputOctets + acctInputGigawords*4*1024*1024*1024
}

func getOutputTotal(form *web.WebForm) int64 {
	var acctOutputOctets = form.GetInt64Val("acctOutputOctets", 0)
	var acctOutputGigawords = form.GetInt64Val("acctOutputGigawords", 0)
	return acctOutputOctets + acctOutputGigawords*4*1024*1024*1024
}

// 更新记账信息
func (m *RadiusManager) UpdateRadiusOnline(form *web.WebForm) error {
	var sessionId = form.GetVal2("acctSessionId", "")
	var statusType = form.GetVal2("acctStatusType", "")
	radOnline := Accounting{
		Username:          form.GetVal("username"),
		NasId:             form.GetVal("nasid"),
		NasAddr:           form.GetVal("nasip"),
		NasPaddr:          form.GetVal("nasip"),
		SessionTimeout:    form.GetIntVal("sessionTimeout", 0),
		FramedIpaddr:      form.GetVal2("framedIPAddress", "0.0.0.0"),
		FramedNetmask:     form.GetVal2("framedIPNetmask", common.NA),
		MacAddr:           form.GetVal2("macAddr", common.NA),
		NasPort:           0,
		NasClass:          common.NA,
		NasPortId:         form.GetVal2("nasPortId", common.NA),
		NasPortType:       0,
		ServiceType:       0,
		AcctSessionId:     sessionId,
		AcctSessionTime:   form.GetIntVal("acctSessionTime", 0),
		AcctInputTotal:    getInputTotal(form),
		AcctOutputTotal:   getOutputTotal(form),
		AcctInputPackets:  form.GetIntVal("acctInputPackets", 0),
		AcctOutputPackets: form.GetIntVal("acctOutputPackets", 0),
		AcctStartTime:     getAcctStartTime(form.GetVal2("acctSessionTime", "0")),
		LastUpdate:       time.Now(),
	}
	switch statusType {
	case "Start", "Update", "Alive", "Interim-Update":
		ocount, _ := m.GetOnlineCountBySessionid(sessionId)
		if ocount == 0 {
			log.Infof("Add radius online %+v", radOnline)
			return m.AddRadiusOnline(radOnline)
		} else {
			log.Infof("Update radius online %+v", radOnline)
			return m.UpdateRadiusOnlineData(radOnline)
		}
	case "Stop":
		log.Infof("Update radius cdr %+v", radOnline)
		_ = m.AddRadiusAccounting(radOnline)
		return m.DeleteRadiusOnline(sessionId)
	}

	return nil
}
