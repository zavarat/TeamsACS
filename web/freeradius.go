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

package web

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/log"
	"github.com/ca17/teamsacs/common/validutil"
	"github.com/ca17/teamsacs/common/web"

	"github.com/labstack/echo/v4"
)

const (
	RadiusAuthSucces        = "success"
	RadiusAuthFailure       = "failure"
	RadiusMaxSessionTimeout = 864000
	RadiusInterimIntelval   = 120
	RadiusAuthlogLevel      = "all"
)

// 添加认证日志
func (h *HttpHandler) AddAuthlog(username string, nasip string, result string, reason string, level string, cast int64) {
	if level != "all" || result != level {
		err := h.GetManager().GetRadiusManager().AddRadiusAuthLog(username, nasip, result, reason, cast)
		if err != nil {
			log.Error(err)
		}
	}
}

// Authorize processing, if the user exists, the password response is sent back for further verification.
//      #  FreeradiusAuthorize/FreeradiusAuthenticate
//      #
//      #  Code   Meaning       Process body  Module code
//      #  404    not found     no            notfound
//      #  410    gone          no            notfound
//      #  403    forbidden     no            userlock
//      #  401    unauthorized  yes           reject
//      #  204    no content    no            ok
//      #  2xx    successful    yes           ok/updated
//      #  5xx    server error  no            fail
//      #  xxx    -             no            invalid
func (h *HttpHandler) FreeradiusAuthorize(c echo.Context) error {
	var start = time.Now()
	username := strings.TrimSpace(c.FormValue("username"))
	nasip := c.FormValue("nasip")

	user, err := h.GetManager().GetSubscribeManager().GetSubscribeByUser(username)
	if err != nil {
		h.AddAuthlog(username, nasip, RadiusAuthFailure, "user query err"+err.Error(), RadiusAuthlogLevel, time.Since(start).Milliseconds())
		return c.JSON(501, echo.Map{"Reply-Message": "user query error, reject auth, " + err.Error()})
	}

	// Check user status
	if user.Status == "disabled" {
		h.AddAuthlog(username, nasip, RadiusAuthFailure, "user disabled", RadiusAuthlogLevel, time.Since(start).Milliseconds())
		return c.JSON(501, echo.Map{"Reply-Message": "user status disabled, reject auth"})
	}

	// Check user expiration
	if user.ExpireTime.Time().Before(time.Now()) {
		h.AddAuthlog(username, nasip, RadiusAuthFailure, "user expire", RadiusAuthlogLevel, time.Since(start).Milliseconds())
		return c.JSON(501, echo.Map{"Reply-Message": "user expire, reject auth"})
	}

	// Evaluation of online limit
	// Current number online
	count, err := h.GetManager().GetRadiusManager().GetOnlineCount(username)
	if err != nil {
		h.AddAuthlog(username, nasip, RadiusAuthFailure, "user query count err"+err.Error(), RadiusAuthlogLevel, time.Since(start).Milliseconds())
		return c.JSON(501, echo.Map{"Reply-Message": "user online count fetch error, reject auth, " + err.Error()})
	}
	if count > 0 && user.Profile.ActiveNum > 0 && count >= int64(user.Profile.ActiveNum) {
		h.AddAuthlog(username, nasip, RadiusAuthFailure, "user online limit", RadiusAuthlogLevel, time.Since(start).Milliseconds())
		return c.JSON(501, echo.Map{"Reply-Message": "user online over limit, reject auth"})
	}

	// freeradius response
	resp := map[string]interface{}{}
	resp["control:Cleartext-Password"] = strings.TrimSpace(user.Password)
	resp["reply:Mikrotik-Rate-Limit"] = fmt.Sprintf("%dk/%dk", user.Profile.UpRate, user.Profile.DownRate)
	sessionTimeout := user.ExpireTime.Time().Sub(time.Now()).Seconds()
	resp["reply:Session-Timeout"] = fmt.Sprintf("%d", int64(sessionTimeout))

	// Set address pool or static IP
	if common.IsNotEmptyAndNA(user.Ipaddr) && validutil.IsIP(user.Ipaddr){
		resp["Framed-IP-Address"] = user.Ipaddr
	} else if common.IsNotEmptyAndNA(user.Profile.AddrPool){
		resp["Framed-Pool"] = user.Profile.AddrPool
	}

	h.AddAuthlog(username, nasip, RadiusAuthSucces, RadiusAuthSucces, RadiusAuthlogLevel, time.Since(start).Milliseconds())

	return c.JSON(http.StatusOK, resp)
}

// Authenticate processing
//     #  FreeradiusAuthorize/FreeradiusAuthenticate
//     #
//     #  Code   Meaning       Process body  Module code
//     #  404    not found     no            notfound
//     #  410    gone          no            notfound
//     #  403    forbidden     no            userlock
//     #  401    unauthorized  yes           reject
//     #  204    no content    no            ok
//     #  2xx    successful    yes           ok/updated
//     #  5xx    server error  no            fail
//     #  xxx    -             no            invalid
func (h *HttpHandler) FreeradiusAuthenticate(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// Postauth processing
func (h *HttpHandler) FreeradiusPostauth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// Accounting processing
func (h *HttpHandler) FreeradiusAccounting(c echo.Context) error {
	webform := web.NewWebForm(c)
	err := h.GetManager().GetRadiusManager().UpdateRadiusOnline(webform)
	if err != nil {
		log.Error(err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{})
}
