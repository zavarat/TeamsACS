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
	"strconv"
	"time"
)

type DataObject map[string]string

func (d DataObject) GetStringValue(key string, defval string) string {
	val, ok := d[key]
	if !ok || val == "" {
		return defval
	}
	return val
}

func (d DataObject) GetInt64Value(key string, defval int64) int64 {
	val, ok := d[key]
	if ok {
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return defval
		}
		return v
	}
	return defval
}

func (d DataObject) GetIntValue(key string, defval int) int {
	val, ok := d[key]
	if ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			return defval
		}
		return v
	}
	return defval
}


func (d DataObject) GetDateValue(key string, defval time.Time) time.Time {
	val, ok := d[key]
	if ok {
		var result = defval
		var err error
		if len(val) == 19 {
			result, err = time.Parse("2006-01-02 15:04:05", val)
		}else{
			result, err = time.Parse("2006-01-02 15:04:05 Z0700 MST", val)
		}
		if err != nil {
			return defval
		}
		return result
	}
	return defval
}


