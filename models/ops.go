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

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ca17/teamsacs/common/log"
)

type OpsManager struct{ *ModelManager }

func (m *ModelManager) GetOpsManager() *OpsManager {
	store, _ := m.ManagerMap.Get("OpsManager")
	return store.(*OpsManager)
}

func (m *OpsManager) AddOpsLog(username string, srcip string, action string, remark string)  {
	authlog := OpsLog{
		Username:  username,
		Srcip:     srcip,
		Action:    action,
		Remark:    remark,
		Timestamp: primitive.NewDateTimeFromTime(time.Now()),
	}
	coll := m.GetTeamsAcsCollection(TeamsacsOpslog)
	_, err := coll.InsertOne(context.TODO(), authlog)
	if err != nil {
		log.Error(err)
	}
}

