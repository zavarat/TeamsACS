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

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4/middleware"
	cmap "github.com/orcaman/concurrent-map"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/gmail"
	"github.com/ca17/teamsacs/common/mongodb"
	"github.com/ca17/teamsacs/common/tpl"
	"github.com/ca17/teamsacs/config"
)

const (
	MDBTeamsacs        = "teamsacs"
	MDBGenieacs        = "genieacs"
	TeamsacsConfig     = "config"
	TeamsacsOperator   = "operator"
	TeamsacsSubscribe  = "subscribe"
	TeamsacsOpslog     = "opslog"
	TeamsacsLdap       = "ldap"
	TeamsacsVpe       = "vpe"
	TeamsacsCpe       = "cpe"
	TeamsacsOnline     = "online"
	TeamsacsProfile    = "profile"
	TeamsacsAccounting = "accounting"
	TeamsacsAuthlog = "authlog"

	GenieacsDevices = "devices"
	GenieacsFaults  = "faults"
	GenieacsTasks   = "tasks"
	GenieacsPresets = "presets"
)
type Doc = map[string]interface{}

type ModelManager struct {
	Config       *config.AppConfig
	Mongo        *mongo.Client
	Sched        *gocron.Scheduler
	TplRender    *tpl.CommonTemplate
	Location     *time.Location
	WebJwtConfig *middleware.JWTConfig
	MailSender   *gmail.MailSender
	ManagerMap   cmap.ConcurrentMap
	Dev          bool
}

func NewModelManager(appconfig *config.AppConfig, dev bool) *ModelManager {
	m := &ModelManager{Config: appconfig, Dev: dev}
	m.ManagerMap = cmap.New()
	_mongodb, err := mongodb.GetMongodbClient(appconfig.Mongodb)
	common.Must(err)
	m.Mongo = _mongodb
	loc, err := time.LoadLocation(appconfig.System.Location)
	common.Must(err)
	m.Location = loc
	m.registerManagers()
	m.TplRender = tpl.NewCommonTemplate([]string{"/resources/templates"}, m.Dev, m.GetTemplateFuncMap())
	go m.StartScheduler()
	return m
}

func (m *ModelManager) registerManagers()  {
	m.ManagerMap.Set("SubscribeManager", &SubscribeManager{m})
	m.ManagerMap.Set("RadiusManager", &RadiusManager{m})
	m.ManagerMap.Set("VpeManager", &VpeManager{m})
	m.ManagerMap.Set("OpsManager", &OpsManager{m})
	m.ManagerMap.Set("CpeManager", &CpeManager{m})
	m.ManagerMap.Set("ConfigManager", &ConfigManager{m})
	m.ManagerMap.Set("GenieacsManager", &GenieacsManager{m})
}


func (m *ModelManager) GetTeamsAcsCollection(coll string) *mongo.Collection {
	return m.Mongo.Database(MDBTeamsacs).Collection(coll)
}

func (m *ModelManager) GetGenieAcsCollection(coll string) *mongo.Collection {
	return m.Mongo.Database(MDBTeamsacs).Collection(coll)
}

func (m *ModelManager) GetTemplateFuncMap() map[string]interface{} {
	return map[string]interface{}{
		"Pagever": func() int64 {
			return time.Now().Unix()
		},
	}
}

