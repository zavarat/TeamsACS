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

package freeradius

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/ca17/teamsacs/config"
	"github.com/ca17/teamsacs/models"
)

type RestResult struct {
	Code    int         `json:"code"`
	Msgtype string      `json:"msgtype"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type WebContext struct {
	Manager *models.ModelManager
	Config  *config.AppConfig
}

// WebHandler
type WebHandler interface {
	InitRouter(group *echo.Group)
}

type HttpHandler struct {
	Ctx *WebContext
}

func NewHttpHandler(ctx *WebContext) HttpHandler {
	return HttpHandler{Ctx: ctx}
}

func (h *HttpHandler) InitRouter(group *echo.Group) {

}

func (h *HttpHandler) GetConfig() *config.AppConfig {
	return h.Ctx.Config
}

func (h *HttpHandler) GetManager() *models.ModelManager {
	return h.Ctx.Manager
}

func (h *HttpHandler) ParseFormInt64(c echo.Context, name string) (int64, error) {
	return strconv.ParseInt(c.FormValue("id"), 10, 64)

}
