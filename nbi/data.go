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
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ca17/teamsacs/common"
)


// A generic data CRUD management API with no predefined schema,
// storing extra data that you may not use at all, but you'll probably use a lot.


// QueryData
func (h *HttpHandler) QueryData(c echo.Context) error {
	params := h.RequestParse(c)
	params["collname"] = c.Param("collname")
	data, err := h.GetManager().GetDataManager().QueryDatas(params)
	common.Must(err)
	return c.JSON(http.StatusOK, data)
}

// QueryData
func (h *HttpHandler) QueryDataOptions(c echo.Context) error {
	params := h.RequestParse(c)
	params["collname"] = c.Param("collname")
	data, err := h.GetManager().GetDataManager().QueryDataOptions(params)
	common.Must(err)
	return c.JSON(http.StatusOK, data)
}

// AddData
func (h *HttpHandler) GetData(c echo.Context) error {
	params := h.RequestParse(c)
	params["collname"] = c.Param("collname")
	r, err := h.GetManager().GetDataManager().GetData(params)
	common.Must(err)
	return c.JSON(http.StatusOK, h.RestResult(r))
}

// AddData
func (h *HttpHandler) AddData(c echo.Context) error {
	params := h.RequestParse(c)
	params["collname"] = c.Param("collname")
	common.Must(h.GetManager().GetDataManager().AddData(params))
	return c.JSON(http.StatusOK, h.RestSucc("Success"))
}

// UpdateData
func (h *HttpHandler) UpdateData(c echo.Context) error {
	params := h.RequestParse(c)
	params["collname"] = c.Param("collname")
	common.Must(h.GetManager().GetDataManager().UpdateData(params))
	return c.JSON(http.StatusOK, h.RestSucc("Success"))
}

// DeleteData
func (h *HttpHandler) DeleteData(c echo.Context) error {
	params := h.RequestParse(c)
	params["collname"] = c.Param("collname")
	common.Must(h.GetManager().GetDataManager().DeleteData(params))
	return c.JSON(http.StatusOK, h.RestSucc("Success"))
}


