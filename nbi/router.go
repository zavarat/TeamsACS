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
	e.Any("/nbi/cpe/delete", h.DeleteCpe)
	e.POST( "/nbi/cpe/add", h.AddCpe)
	e.POST( "/nbi/cpe/update", h.UpdateCpe)

	// Vpe apis
	e.Any("/nbi/vpe/query", h.QueryVpe)
	e.Any("/nbi/vpe/delete", h.DeleteVpe)
	e.POST( "/nbi/vpe/add", h.AddVpe)
	e.POST( "/nbi/vpe/update", h.UpdateVpe)

	// Subscribe apis
	e.Any("/nbi/subscribe/query", h.QuerySubscribe)
	e.Any("/nbi/subscribe/delete", h.DeleteSubscribe)
	e.POST("/nbi/subscribe/add",h.AddSubscribe)
	e.POST("/nbi/subscribe/update", h.UpdateSubscribe)

	// opr apis
	e.Any("/nbi/opr/query", h.QueryOperator)
	e.Any("/nbi/opr/delete", h.DeleteOperator)
	e.POST( "/nbi/opr/add",h.AddOperator)
	e.POST( "/nbi/opr/update", h.UpdateOperator)

	// opr apis
	e.Any("/nbi/data/:collname/query", h.QueryData)
	e.Any("/nbi/data/:collname/options", h.QueryDataOptions)
	e.Any("/nbi/data/:collname/get", h.GetData)
	e.Any("/nbi/data/:collname/delete", h.DeleteData)
	e.POST( "/nbi/data/:collname/add",h.AddData)
	e.POST( "/nbi/data/:collname/update", h.UpdateData)

	// radius apis
	e.Any("/nbi/radius/accounting/query", h.QueryRadiusAccounting)
	e.Any("/nbi/radius/authlog/query", h.QueryRadiusAuthlog)
	e.Any("/nbi/radius/online/query", h.QueryRadiusOnline)

	// config apis
	e.POST("/nbi/config/radius/update", h.UpdateRadiusConfigs)
	e.POST("/nbi/config/update", h.UpdateConfig)
	e.Any("/nbi/config/query", h.QueryConfig)

	// token
	e.POST( "/nbi/token", h.RequestToken)
}
