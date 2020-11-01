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
)

func (h *HttpHandler) InitAllRouter(e *echo.Echo) {
	// mikrotik cpe query apis
	e.Any("/nbi/mikrotik/device/interfaces", h.QueryMikrotikDeviceInterfaces)
	e.Any("/nbi/mikrotik/device/pppinterfaces", h.QueryMikrotikDevicePPPInterfaces)
	e.Any("/nbi/mikrotik/device/ipinterfaces", h.QueryMikrotikDeviceIpInterfaces)
	e.Any("/nbi/mikrotik/device/routers", h.QueryMikrotikDeviceRouters)
	e.Any("/nbi/mikrotik/device/dns", h.QueryMikrotikDeviceDnsClientServer)

	// Cpe apis
	e.Any("/nbi/cpe/query", h.QueryCpe)
	e.Add(http.MethodPost, "/nbi/cpe/add", h.AddCpe)
	e.Add(http.MethodPost, "/nbi/cpe/update", h.UpdateCpe)
	e.Any("/nbi/cpe/delete", h.DeleteCpe)

	// Vpe apis
	e.Any("/nbi/vpe/query", h.QueryVpe)
	e.Add(http.MethodPost, "/nbi/vpe/add", h.AddVpe)
	e.Add(http.MethodPost, "/nbi/vpe/update", h.UpdateVpe)
	e.Any("/nbi/vpe/delete", h.DeleteVpe)

	// Subscribe apis
	e.Any("/nbi/subscribe/query", h.QuerySubscribe)
	e.Add(http.MethodPost, "/nbi/subscribe/add",h.AddSubscribe)
	e.Add(http.MethodPost, "/nbi/subscribe/update", h.UpdateSubscribe)
	e.Any("/nbi/subscribe/delete", h.DeleteSubscribe)

	// opr apis
	e.Any("/nbi/opr/query", h.QueryOperator)
	e.Add(http.MethodPost, "/nbi/opr/add",h.AddOperator)
	e.Add(http.MethodPost, "/nbi/opr/update", h.UpdateOperator)
	e.Any("/nbi/opr/delete", h.DeleteOperator)

	// radius apis
	e.Any("/nbi/radius/accounting/query", h.QueryRadiusAccounting)
	e.Any("/nbi/radius/authlog/query", h.QueryRadiusAuthlog)
	e.Any("/nbi/radius/online/query", h.QueryRadiusOnline)

	// token
	e.Add(http.MethodPost, "/nbi/token", h.RequestToken)
}
