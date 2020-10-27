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

package config

import (
	"gopkg.in/yaml.v2"
	"testing"
)

func TestEncodeYaml(t *testing.T) {
	r, _ := yaml.Marshal(DefaultAppConfig)
	t.Log(string(r))
}

func TestDecodeYaml(t *testing.T) {
	conf := &AppConfig{}
	ystr := `system:
  appid: TeamsACS
  workdir: /var/teamsacs
  syslog_addr: ""
  debug: false
web:
  host: 0.0.0.0
  port: 18998
  debug: true
  jwt_secret: 9b6de5cc-0738-4bf1-acs1-0f568ac9da37
mongodb:
  url: mongodb://127.0.0.1:27017
  user: ""
  passwd: ""
grpc:
  host: 0.0.0.0
  port: 18999
  debug: true
radiusd:
  host: 0.0.0.0
  auth_port: 1812
  acct_port: 1813
  debug: true`
	yaml.Unmarshal([]byte(ystr), conf)

	t.Log(conf)
}
