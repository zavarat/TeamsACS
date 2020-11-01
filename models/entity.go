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
)

type Attributes = map[string]interface{}

type Config struct {
	ID     string `bson:"_id,omitempty" json:"id"`
	Type   string `bson:"type" json:"type"`
	Name   string `bson:"name" json:"name"`
	Value  string `bson:"value" json:"value"`
	Remark string `bson:"remark" json:"remark"`
}

type Operator struct {
	ID        string   `bson:"_id,omitempty" json:"id"`
	Email     string   `bson:"email,omitempty" json:"email"`
	Username  string   `bson:"username,omitempty" json:"username"`
	Level  string   `bson:"level" json:"level"`
	ApiSecret string   `bson:"api_secret,omitempty" json:"api_secret"`
	Status    string   `bson:"status,omitempty" json:"status"`
	Remark    string   `bson:"remark,omitempty" json:"remark"`
}

type OpsLog struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Username  string    `bson:"username,omitempty" json:"username"`
	Srcip     string    `bson:"srcip,omitempty" json:"srcip"`
	Action    string    `bson:"action,omitempty" json:"action"`
	Remark    string    `bson:"remark,omitempty" json:"remark"`
	Timestamp time.Time `bson:"timestamp,omitempty" json:"timestamp"`
}

// VariableConfig
type VariableConfig struct {
	ID     string     `bson:"_id,omitempty" json:"id"`
	Vendor string     `bson:"vendor" json:"vendor"`
	Group  string     `bson:"group" json:"group"`
	Values Attributes `bson:"values" json:"values"`
	Remark string     `bson:"remark" json:"remark"`
}

// Cpe
// attrs: Extended Attributes
type Cpe struct {
	Id         string     `bson:"_id,omitempty" json:"id"`
	Sn         string     `bson:"sn" json:"sn"`
	DeviceId   string     `bson:"device_id" json:"device_id" `
	Attrs      Attributes `bson:"attrs" json:"attrs" `
	CreateTime time.Time  `bson:"create_time" json:"create_time" `
	UpdateTime time.Time  `bson:"update_time" json:"update_time" `
}

type Vpe struct {
	ID         string     `bson:"_id,omitempty" json:"id"`
	Sn         string     `bson:"sn" json:"sn"`
	DeviceId   string     `bson:"device_id" json:"device_id" `
	Attrs      Attributes `bson:"attrs" json:"attrs" `
	Identifier string     `bson:"identifier,omitempty" json:"identifier"`
	Name       string     `bson:"name,omitempty" json:"name"`
	Ipaddr     string     `bson:"ipaddr,omitempty" json:"ipaddr"`
	Secret     string     `bson:"secret,omitempty" json:"-"`
	VendorCode string     `bson:"vendor_code,omitempty" json:"vendor_code"`
	CoaPort    int        `bson:"coa_port,omitempty" json:"coa_port"`
	Status     string     `bson:"status,omitempty" json:"status"`
	LdapId     string     `bson:"ldap_id,omitempty" json:"ldap_id,string"`
	Remark     string     `bson:"remark,omitempty" json:"remark"`
}

type Ldap struct {
	ID         string `bson:"_id,omitempty" json:"id"`
	Name       string `bson:"name,omitempty" json:"name"`
	Address    string `bson:"address,omitempty" json:"address"`
	Password   string `bson:"password,omitempty" json:"-"`
	Searchdn   string `bson:"searchdn,omitempty" json:"searchdn"`
	Basedn     string `bson:"basedn,omitempty" json:"basedn"`
	UserFilter string `bson:"user_filter,omitempty" json:"user_filter"`
	Istls      string `bson:"istls,omitempty" json:"istls"`
	Status     string `bson:"status,omitempty" json:"status"`
	Remark     string `bson:"remark,omitempty" json:"remark"`
}

type ProfileAttr struct {
	Domain          string `bson:"domain" json:"domain"`
	InterimInterval int    `bson:"interim_interval" json:"interim_interval"`
	AddrPool        string `bson:"addr_pool" json:"addr_pool"`
	ActiveNum       int    `bson:"active_num" json:"active_num"`
	MfaStatus       string `bson:"mfa_status" json:"mfa_status"`
	UpRate          int    `bson:"up_rate" json:"up_rate"`
	DownRate        int    `bson:"down_rate" json:"down_rate"`
	LimitPolicy     string `bson:"limit_policy" json:"limit_policy"`
	UpLimitPolicy   string `bson:"up_limit_policy" json:"up_limit_policy"`
	DownLimitPolicy string `bson:"down_limit_policy" json:"down_limit_policy"`
}

type Profile struct {
	ProfileAttr
	ID           string `bson:"_id,omitempty" json:"id"`
	Name         string `bson:"name" json:"name"`
	BillTimes    int    `bson:"bill_times" json:"bill_times"`
	BillTimeunit string `bson:"bill_timeunit" json:"bill_timeunit"`
	Status       string `bson:"status" json:"status"`
	Remark       string `bson:"remark" json:"remark"`
}

type Subscribe struct {
	ID         string      `bson:"_id,omitempty" json:"id"`
	VpeSids    []string    `bson:"vpe_sids,omitempty" json:"vpe_sids"`
	Profile    ProfileAttr `bson:"profile,omitempty" json:"profile,omitempty"`
	Realname   string      `bson:"realname,omitempty" json:"realname"`
	Email      string      `bson:"email,omitempty" json:"email"`
	Username   string      `bson:"username,omitempty" json:"username"`
	Password   string      `bson:"password,omitempty" json:"password"`
	MfaStatus  string      `bson:"mfa_status,omitempty" json:"mfa_status"`
	MfaSecret  string      `bson:"mfa_secret,omitempty" json:"mfa_secret"`
	Ipaddr     string      `bson:"ipaddr,omitempty" json:"ipaddr"`
	Macaddr    string      `bson:"macaddr,omitempty" json:"macaddr"`
	Vlanid1    int         `bson:"vlanid_1,omitempty" json:"vlanid1"`
	Vlanid2    int         `bson:"vlanid_2,omitempty" json:"vlanid2"`
	BindMac    int         `bson:"bind_mac,omitempty" json:"bind_mac"`
	BindVlan   int         `bson:"bind_vlan,omitempty" json:"bind_vlan"`
	Status     string      `bson:"status,omitempty" json:"status"`
	Remark     string      `bson:"remark,omitempty" json:"remark"`
	ExpireTime time.Time   `bson:"expire_time,omitempty" json:"expire_time"`
	Timestamp  time.Time   `bson:"timestamp,omitempty" json:"timestamp"`
}

type Accounting struct {
	ID                string    `bson:"_id,omitempty" json:"id"`
	Username          string    `bson:"username,omitempty" json:"username"`
	NasId             string    `bson:"nas_id,omitempty" json:"nas_id"`
	NasAddr           string    `bson:"nas_addr,omitempty" json:"nas_addr"`
	NasPaddr          string    `bson:"nas_paddr,omitempty" json:"nas_paddr"`
	SessionTimeout    int       `bson:"session_timeout,omitempty" json:"session_timeout"`
	FramedIpaddr      string    `bson:"framed_ipaddr,omitempty" json:"framed_ipaddr"`
	FramedNetmask     string    `bson:"framed_netmask,omitempty" json:"framed_netmask"`
	MacAddr           string    `bson:"mac_addr,omitempty" json:"mac_addr"`
	NasPort           int64     `bson:"nas_port,omitempty" json:"nas_port,string"`
	NasClass          string    `bson:"nas_class,omitempty" json:"nas_class"`
	NasPortId         string    `bson:"nas_port_id,omitempty" json:"nas_port_id"`
	NasPortType       int       `bson:"nas_port_type,omitempty" json:"nas_port_type"`
	ServiceType       int       `bson:"service_type,omitempty" json:"service_type"`
	AcctSessionId     string    `bson:"acct_session_id,omitempty" json:"acct_session_id"`
	AcctSessionTime   int       `bson:"acct_session_time,omitempty" json:"acct_session_time"`
	AcctInputTotal    int64     `bson:"acct_input_total,omitempty" json:"acct_input_total,string"`
	AcctOutputTotal   int64     `bson:"acct_output_total,omitempty" json:"acct_output_total,string"`
	AcctInputPackets  int       `bson:"acct_input_packets,omitempty" json:"acct_input_packets"`
	AcctOutputPackets int       `bson:"acct_output_packets,omitempty" json:"acct_output_packets"`
	AcctStartTime     time.Time `bson:"acct_start_time,omitempty" json:"acct_start_time"`
	LastUpdate        time.Time `bson:"last_update,omitempty" json:"last_update"`
	AcctStopTime      time.Time `bson:"acct_stop_time,omitempty" json:"acct_stop_time"`
}

type Authlog struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Username  string    `bson:"username,omitempty" json:"username"`
	NasAddr   string    `bson:"nas_addr,omitempty" json:"nas_addr"`
	Cast      int       `bson:"cast,omitempty" json:"cast"`
	Result    string    `bson:"result,omitempty" json:"result"`
	Reason    string    `bson:"reason,omitempty" json:"reason"`
	Timestamp time.Time `bson:"timestamp,omitempty" json:"timestamp"`
}

// AuthorizationProfile Method

func (a Subscribe) GetExpireTime() time.Time {
	return a.ExpireTime
}

func (a Subscribe) GetInterimInterval() int {
	return a.Profile.InterimInterval
}

func (a Subscribe) GetAddrPool() string {
	return a.Profile.AddrPool
}

func (a Subscribe) GetIpaddr() string {
	return a.Ipaddr
}

func (a Subscribe) GetUpRateKbps() int {
	return a.Profile.UpRate
}

func (a Subscribe) GetDownRateKbps() int {
	return a.Profile.DownRate
}

func (a Subscribe) GetDomain() string {
	return a.Profile.Domain
}

func (a Subscribe) GetLimitPolicy() string {
	return a.Profile.LimitPolicy
}

func (a Subscribe) GetUpLimitPolicy() string {
	return a.Profile.UpLimitPolicy
}

func (a Subscribe) GetDownLimitPolicy() string {
	return a.Profile.DownLimitPolicy
}
