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
	"fmt"

	"github.com/ca17/teamsacs/common"
)

func (a *Subscribe) AddValidate() error {
	if common.IsEmptyOrNA(a.Username) {
		return fmt.Errorf("invalid username")
	}
	if common.IsEmptyOrNA(a.Password) {
		return fmt.Errorf("invalid password")
	}
	return nil
}

func (a *Subscribe) UpdateValidate() error {
	if common.IsEmptyOrNA(a.Username) {
		return fmt.Errorf("invalid username")
	}
	return nil
}

func (a *Vpe) AddValidate() error {
	switch {
	case common.IsEmptyOrNA(a.Sn):
		return fmt.Errorf("invalid sn")
	case common.IsEmptyOrNA(a.Identifier):
		return fmt.Errorf("invalid identifier")
	case common.IsEmptyOrNA(a.Ipaddr):
		return fmt.Errorf("invalid ipaddr")
	case common.IsEmptyOrNA(a.Secret):
		return fmt.Errorf("invalid secret")
	case common.IsEmptyOrNA(a.VendorCode):
		return fmt.Errorf("invalid vendor_code")
	}
	return nil
}

func (a *Operator) AddValidate() error {
	switch {
	case common.IsEmptyOrNA(a.Username):
		return fmt.Errorf("invalid username")
	case common.IsEmptyOrNA(a.Level):
		return fmt.Errorf("invalid level")
	}
	return nil
}
