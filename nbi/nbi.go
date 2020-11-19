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
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/web"
	"github.com/360EntSecGroup-Skylar/excelize"
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
		return c.Render(http.StatusInternalServerError, "err500", map[string]string{"message": err.(error).Error()})
	case string:
		return c.Render(http.StatusInternalServerError, "err500", map[string]string{"message": err.(string)})
	}
	return c.Render(http.StatusInternalServerError, "err500", map[string]string{"message": err.(string)})
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

func (h *HttpHandler) JsonBodyParse(c echo.Context) (web.RequestParams, error) {
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
	querymap := make(map[string]interface{})
	equalmap := make(map[string]interface{})
	filtermap := make(map[string]interface{})
	sortmap := make(map[string]interface{})
	for k, vs := range c.QueryParams() {
		if common.InSlice(k, []string{"start", "count"}){
			query[k] = vs[0]
		} else if strings.HasPrefix(k, "filter[") && vs[0] != ""{
			filtermap[k[7:len(k)-1]] = vs[0]
		}else if strings.HasPrefix(k, "equal[") && vs[0] != ""{
			equalmap[k[6:len(k)-1]] = vs[0]
		} else if strings.HasPrefix(k, "sort[") && vs[0] != ""{
			sortmap[k[5:len(k)-1]] = vs[0]
		}else if vs[0] != "" {
			querymap[k] = vs[0]
		}
	}
	query["querymap"] = querymap
	query["filtermap"] = filtermap
	query["equalmap"] = equalmap
	query["sortmap"] = sortmap
	return query
}

func (h *HttpHandler) RequestParseForm(c echo.Context) web.RequestParams {
	params := make(web.RequestParams)
	data := make(map[string]interface{})
	posts, _ := c.FormParams()
	for k, vs := range posts {
		data[k] = vs[0]
	}
	params["data"] = data
	return params
}

func (h *HttpHandler) RequestParse(c echo.Context) web.RequestParams {
	var params = web.EmptyRequestParams
	var err error
	switch c.Request().Method {
	case http.MethodGet:
		params = h.RequestParseGet(c)
	case http.MethodPost, http.MethodPut:
		ctype := c.Request().Header.Get("Content-Type")
		if ctype == "application/json" {
			params, err = h.JsonBodyParse(c)
			common.Must(err)
		}else if ctype == "application/x-www-form-urlencoded" {
			params = h.RequestParseForm(c)
		}
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


func (h *HttpHandler) FetchExcelData(c echo.Context, sheet string) ([]map[string]string, error) {
	file, err := c.FormFile("upload")
	if err != nil {
		return nil, err
	}
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	f, err := excelize.OpenReader(src)
	if err != nil {
		return nil, err
	}
	// Get all cells on Sheet1
	rows := f.GetRows(sheet)
	head := make(map[int]string)
	var data []map[string]string
	for i, row := range rows {
		item := make(map[string]string)
		for k, colCell := range row {
			if i == 0 {
				head[k] = colCell
			} else {
				item[head[k]] = colCell
			}
		}
		if i == 0 {
			continue
		}
		data = append(data, item)
	}

	return data, nil
}
