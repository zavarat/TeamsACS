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


type CollTemplate struct {
	ID    string     `bson:"_id,omitempty" json:"id,omitempty"`
	Attrs Attributes `bson:"attrs" json:"attrs,omitempty"`
}

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

