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
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ca17/teamsacs/common/web"
)

type QueryForm struct {
	Sort    string `query:"sort" form:"sort"`
	Size    int64 `query:"size" form:"size"`
}

type SubscribeQueryForm struct {
	QueryForm
	Username string `query:"username" form:"username"`
}


type QueryResult = []map[string]interface{}

func (m *ModelManager) QueryItems(form *web.WebForm, collatiion string) (QueryResult, error) {
	var findOptions = options.Find()
	coll := m.GetTeamsAcsCollection(collatiion)
	var q = bson.M{}
	for qname, vals := range form.Gets {
		q[qname] = vals[0]
	}
	cur, err := coll.Find(context.TODO(), q, findOptions)
	if err != nil {
		return nil, err
	}
	items := make(QueryResult, 0)
	for cur.Next(context.TODO()) {
		var elem map[string]interface{}
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		} else {
			items = append(items, elem)
		}
	}
	return items, nil
}

func (m *ModelManager) QueryPagerItems(params web.RequestParams, collatiion string) (*web.PageResult, error) {
	var findOptions = options.Find()
	var pos = params.GetInt64WithDefval("start", 0)
	findOptions.SetSkip(pos)
	findOptions.SetLimit(params.GetInt64WithDefval("count", 40))
	coll := m.GetTeamsAcsCollection(collatiion)
	var q = bson.M{}
	for qname, val := range params.GetParamMap("querymap") {
		q[qname] = val
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
