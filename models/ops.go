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

type OpsManager struct{ *ModelManager }

func (m *ModelManager) GetOpsManager() *OpsManager {
	store, _ := m.ManagerMap.Get("OpsManager")
	return store.(*OpsManager)
}


// ExistOperator
func (m *OpsManager) ExistOperator(username string) bool {
	coll := m.GetTeamsAcsCollection(TeamsacsOperator)
	count, _ := coll.CountDocuments(context.TODO(), bson.M{"username": username})
	return count > 0
}

// QueryOperators
func (m *OpsManager) QueryOperators(params web.RequestParams) (*web.PageResult, error) {
	return m.QueryPagerItems(params, TeamsacsOperator)
}

// GetOperator
func (m *OpsManager) GetOperator(username string) (*Operator, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsOperator)
	doc := coll.FindOne(context.TODO(), bson.M{"username": username})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Operator)
	err = doc.Decode(result)
	return result, err
}


// UpdateSubscribe
func (m *OpsManager) UpdateApiSecret(username string) (string, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsOperator)
	apisecret := common.UUID()
	_, err := coll.UpdateOne(context.TODO(), bson.M{"username": username}, bson.M{"$set":bson.M{"api_secret":apisecret}})
	return apisecret, err
}


// UpdateOperator
// update by username
func (m *OpsManager) UpdateOperator(operator *Operator) error {
	coll := m.GetTeamsAcsCollection(TeamsacsOperator)
	query := bson.M{"username": operator.Username}
	data := bson.M{}
	if common.InSlice(operator.Level, []string{constant.NBIAdminLevel,constant.NBIOprLevel}) {
		data["level"] = operator.Level
	}
	if common.InSlice(operator.Status, []string{constant.ENABLED,constant.DISABLED}) {
		data["status"] = operator.Status
	}
	data["email"] = operator.Email
	data["remark"] = operator.Remark

	update := bson.M{"$set": data}
	_, err := coll.UpdateOne(context.TODO(), query, update)
	return err
}

// AddOperator
func (m *OpsManager) AddOperator(operator *Operator) (string, error) {
	if err := operator.AddValidate(); err != nil {
		return "", err
	}
	if m.ExistOperator(operator.Username) {
		return "", fmt.Errorf("operator exists")
	}
	if common.IsEmptyOrNA(operator.ApiSecret){
		operator.ApiSecret = common.UUID()
	}
	operator.Status = constant.ENABLED
	r, err := m.GetTeamsAcsCollection(TeamsacsOperator).InsertOne(context.TODO(), operator)
	if err != nil {
		return "", err
	}
	return r.InsertedID.(string), err
}

// DeleteSubscribe
func (m *OpsManager) DeleteOperator(username string) error {
	if common.IsEmptyOrNA(username) {
		return fmt.Errorf("username is empty or NA")
	}
	_, err := m.GetTeamsAcsCollection(TeamsacsOperator).DeleteOne(context.TODO(), bson.M{"username": username})
	return err
}
