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
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/models"
)

// QueryVpe
func (h *HttpHandler) QueryVpe(c echo.Context) error {
	var result = make(map[string]interface{})
	params := h.RequestParse(c)
	data, err := h.GetManager().GetVpeManager().QueryVpes(params)
	if err != nil {
		return h.GetInternalError(err)
	}
	result["data"] = data
	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) AddVpe(c echo.Context) error {
	item := new(models.Vpe)
	common.Must(c.Bind(item))
	if common.IsEmptyOrNA(item.Sn) {
		common.Must(fmt.Errorf("sn is empty or NA"))
	}
	common.Must(h.GetManager().GetVpeManager().AddVpe(item))
	h.AddOpsLog(c, fmt.Sprintf("Add Vpe sn=%s", c.FormValue("sn")))
	return c.JSON(200, h.RestSucc("Success"))
}

func (h *HttpHandler) UpdateVpe(c echo.Context) error {
	item := new(models.Vpe)
	common.Must(c.Bind(item))
	err := h.GetManager().GetVpeManager().UpdateVpe(item)
	common.Must(err)
	h.AddOpsLog(c, fmt.Sprintf("Update Vpe sn=%s", item.Sn))
	return c.JSON(http.StatusOK, h.RestSucc("Success"))
}

func (h *HttpHandler) DeleteVpe(c echo.Context) error {
	params := h.RequestParse(c)
	sn := params.GetMustString("sn")
	common.Must(h.GetManager().GetVpeManager().DeleteVpe(sn))
	h.AddOpsLog(c, fmt.Sprintf("Delete Vpe sn=%s", sn))
	return c.JSON(http.StatusOK, h.RestSucc("Success"))
}

