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

package nbi

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/web"
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

func (h *HttpHandler) GetInternalError(err interface{}) *echo.HTTPError {
	switch err.(type) {
	case error:
		return echo.NewHTTPError(http.StatusInternalServerError, err.(error).Error())
	case string:
		return echo.NewHTTPError(http.StatusInternalServerError, err.(string))
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err)
}

func (h *HttpHandler) GoInternalErrPage(c echo.Context, err interface{}) error {
	switch err.(type) {
	case error:
		return c.Render(http.StatusInternalServerError, "err500", map[string]string{"message":err.(error).Error()})
	case string:
		return c.Render(http.StatusInternalServerError, "err500", map[string]string{"message":err.(string)})
	}
	return c.Render(http.StatusInternalServerError, "err500", map[string]string{"message":err.(string)})
}


func (h *HttpHandler) RestResult(data interface{}) *RestResult {
	return &RestResult{
		Code:    0,
		Msgtype: "info",
		Msg:     "Operation Success",
		Data:    data,
	}
}

func (h *HttpHandler) RestSucc(msg string) *RestResult {
	return &RestResult{
		Code:    0,
		Msgtype: "info",
		Msg:     msg,
	}
}

func (h *HttpHandler) RestError(msg string) *RestResult {
	return &RestResult{
		Code:    9999,
		Msgtype: "error",
		Msg:     msg,
	}
}

func (h *HttpHandler) ParseFormInt64(c echo.Context, name string) (int64, error) {
	return strconv.ParseInt(c.FormValue("id"), 10, 64)

}

func (h *HttpHandler) GetJwtData(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}

// Get current api user name
func (h *HttpHandler) GetUsername(c echo.Context) string {
	jd := h.GetJwtData(c)
	username, _ := jd["usr"]
	return username.(string)
}

// Get current api user level
func (h *HttpHandler) GetUserLevel(c echo.Context) string {
	jd := h.GetJwtData(c)
	level, _ := jd["lvl"]
	return level.(string)
}

// Get current api user id
func (h *HttpHandler) GetUserId(c echo.Context) string {
	jd := h.GetJwtData(c)
	uid, _ := jd["uid"]
	return uid.(string)
}

// Adding operational logs for audit
func (h *HttpHandler) AddOpsLog(c echo.Context, desc string)  {
	jd := h.GetJwtData(c)
	h.GetManager().GetOpsManager().AddOpsLog(jd["usr"].(string), c.RealIP(), c.Path(), html.EscapeString(desc))
}

func (h *HttpHandler) JsonBodyParse(c echo.Context) (web.RequestParams, error)  {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return nil, err
	}
	query := make(web.RequestParams)
	err = json.Unmarshal(body, &query)
	return query, err
}

func (h *HttpHandler) RequestParseGet(c echo.Context) web.RequestParams {
	query := make(web.RequestParams)
	for k, vs := range c.QueryParams() {
		query[k] = vs[0]
	}
	return query
}


func (h *HttpHandler) RequestParse(c echo.Context) web.RequestParams {
	var params = web.EmptyRequestParams
	var err error
	switch c.Request().Method {
	case http.MethodGet:
		params = h.RequestParseGet(c)
	case http.MethodPost, http.MethodPut:
		params, err = h.JsonBodyParse(c)
		common.Must(err)
	}
	return params
}




type HTTPError struct {
	Code     int         `json:"-"`
	Message  interface{} `json:"message"`
	Internal error       `json:"-"` // Stores the error returned by an external dependency
}

func NewHTTPError(code int, message ...interface{}) *HTTPError {
	he := &HTTPError{Code: code, Message: http.StatusText(code)}
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d:%s", e.Code, e.Message)
}
