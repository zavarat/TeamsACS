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
	ID     string `bson:"_id,omitempty" json:"id,omitempty"`
	Type   string `bson:"type" json:"type,omitempty"`
	Name   string `bson:"name" json:"name,omitempty"`
	Value  string `bson:"value" json:"value,omitempty"`
	Remark string `bson:"remark" json:"remark,omitempty"`
}

type Operator struct {
	ID        string   `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string   `bson:"email,omitempty" json:"email,omitempty"`
	Username  string   `bson:"username,omitempty" json:"username,omitempty"`
	Level  string   `bson:"level" json:"level,omitempty"`
	ApiSecret string   `bson:"api_secret,omitempty" json:"api_secret,omitempty"`
	Status    string   `bson:"status,omitempty" json:"status,omitempty"`
	Remark    string   `bson:"remark,omitempty" json:"remark,omitempty"`
}

type OpsLog struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string    `bson:"username,omitempty" json:"username,omitempty"`
	Srcip     string    `bson:"srcip,omitempty" json:"srcip,omitempty"`
	Action    string    `bson:"action,omitempty" json:"action,omitempty"`
	Remark    string    `bson:"remark,omitempty" json:"remark,omitempty"`
	Timestamp time.Time `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
}

// VariableConfig
type VariableConfig struct {
	ID     string     `bson:"_id,omitempty" json:"id,omitempty"`
	Vendor string     `bson:"vendor" json:"vendor,omitempty"`
	Group  string     `bson:"group" json:"group,omitempty"`
	Values Attributes `bson:"values" json:"values,omitempty"`
	Remark string     `bson:"remark" json:"remark,omitempty"`
}

// Cpe
// attrs: Extended Attributes
type Cpe struct {
	Id         string     `bson:"_id,omitempty" json:"id,omitempty"`
	Sn         string     `bson:"sn" json:"sn,omitempty"`
	DeviceId   string     `bson:"device_id" json:"device_id,omitempty" `
	Attrs      Attributes `bson:"attrs" json:"attrs,omitempty" `
	CreateTime time.Time  `bson:"create_time" json:"create_time,omitempty" `
	UpdateTime time.Time  `bson:"update_time" json:"update_time,omitempty" `
}

type Vpe struct {
	ID         string     `bson:"_id,omitempty" json:"id,omitempty"`
	Sn         string     `bson:"sn" json:"sn,omitempty"`
	DeviceId   string     `bson:"device_id" json:"device_id,omitempty" `
	Attrs      Attributes `bson:"attrs" json:"attrs,omitempty" `
	Identifier string     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name       string     `bson:"name,omitempty" json:"name,omitempty"`
	Ipaddr     string     `bson:"ipaddr,omitempty" json:"ipaddr,omitempty"`
	Secret     string     `bson:"secret,omitempty" json:"secret,omitempty"`
	VendorCode string     `bson:"vendor_code,omitempty" json:"vendor_code,omitempty"`
	CoaPort    int        `bson:"coa_port,omitempty" json:"coa_port,omitempty"`
	Status     string     `bson:"status,omitempty" json:"status,omitempty"`
	LdapId     string     `bson:"ldap_id,omitempty" json:"ldap_id,omitempty,string"`
	Remark     string     `bson:"remark,omitempty" json:"remark,omitempty"`
}

type Ldap struct {
	ID         string `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string `bson:"name,omitempty" json:"name,omitempty"`
	Address    string `bson:"address,omitempty" json:"address,omitempty"`
	Password   string `bson:"password,omitempty" json:"-"`
	Searchdn   string `bson:"searchdn,omitempty" json:"searchdn,omitempty"`
	Basedn     string `bson:"basedn,omitempty" json:"basedn,omitempty"`
	UserFilter string `bson:"user_filter,omitempty" json:"user_filter,omitempty"`
	Istls      string `bson:"istls,omitempty" json:"istls,omitempty"`
	Status     string `bson:"status,omitempty" json:"status,omitempty"`
	Remark     string `bson:"remark,omitempty" json:"remark,omitempty"`
}

type ProfileAttr struct {
	Domain          string `bson:"domain" json:"domain,omitempty"`
	InterimInterval int    `bson:"interim_interval" json:"interim_interval,omitempty"`
	AddrPool        string `bson:"addr_pool" json:"addr_pool,omitempty"`
	ActiveNum       int    `bson:"active_num" json:"active_num,omitempty"`
	MfaStatus       string `bson:"mfa_status" json:"mfa_status,omitempty"`
	UpRate          int    `bson:"up_rate" json:"up_rate,omitempty"`
	DownRate        int    `bson:"down_rate" json:"down_rate,omitempty"`
	LimitPolicy     string `bson:"limit_policy" json:"limit_policy,omitempty"`
	UpLimitPolicy   string `bson:"up_limit_policy" json:"up_limit_policy,omitempty"`
	DownLimitPolicy string `bson:"down_limit_policy" json:"down_limit_policy,omitempty"`
}

type Profile struct {
	ProfileAttr
	ID           string `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string `bson:"name" json:"name,omitempty"`
	BillTimes    int    `bson:"bill_times" json:"bill_times,omitempty"`
	BillTimeunit string `bson:"bill_timeunit" json:"bill_timeunit,omitempty"`
	Status       string `bson:"status" json:"status,omitempty"`
	Remark       string `bson:"remark" json:"remark,omitempty"`
}

type Subscribe struct {
	ID         string      `bson:"_id,omitempty" json:"id,omitempty"`
	VpeSids    []string    `bson:"vpe_sids,omitempty" json:"vpe_sids,omitempty"`
	Profile    ProfileAttr `bson:"profile,omitempty" json:"profile,omitempty,omitempty"`
	Realname   string      `bson:"realname,omitempty" json:"realname,omitempty"`
	Email      string      `bson:"email,omitempty" json:"email,omitempty"`
	Username   string      `bson:"username,omitempty" json:"username,omitempty"`
	Password   string      `bson:"password,omitempty" json:"password,omitempty"`
	MfaStatus  string      `bson:"mfa_status,omitempty" json:"mfa_status,omitempty"`
	MfaSecret  string      `bson:"mfa_secret,omitempty" json:"mfa_secret,omitempty"`
	Ipaddr     string      `bson:"ipaddr,omitempty" json:"ipaddr,omitempty"`
	Macaddr    string      `bson:"macaddr,omitempty" json:"macaddr,omitempty"`
	Vlanid1    int         `bson:"vlanid_1,omitempty" json:"vlanid1,omitempty"`
	Vlanid2    int         `bson:"vlanid_2,omitempty" json:"vlanid2,omitempty"`
	BindMac    int         `bson:"bind_mac,omitempty" json:"bind_mac,omitempty"`
	BindVlan   int         `bson:"bind_vlan,omitempty" json:"bind_vlan,omitempty"`
	Status     string      `bson:"status,omitempty" json:"status,omitempty"`
	Remark     string      `bson:"remark,omitempty" json:"remark,omitempty"`
	ExpireTime time.Time   `bson:"expire_time,omitempty" json:"expire_time,omitempty"`
	Timestamp  time.Time   `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
}

type Accounting struct {
	ID                string    `bson:"_id,omitempty" json:"id,omitempty"`
	Username          string    `bson:"username,omitempty" json:"username,omitempty"`
	NasId             string    `bson:"nas_id,omitempty" json:"nas_id,omitempty"`
	NasAddr           string    `bson:"nas_addr,omitempty" json:"nas_addr,omitempty"`
	NasPaddr          string    `bson:"nas_paddr,omitempty" json:"nas_paddr,omitempty"`
	SessionTimeout    int       `bson:"session_timeout,omitempty" json:"session_timeout,omitempty"`
	FramedIpaddr      string    `bson:"framed_ipaddr,omitempty" json:"framed_ipaddr,omitempty"`
	FramedNetmask     string    `bson:"framed_netmask,omitempty" json:"framed_netmask,omitempty"`
	MacAddr           string    `bson:"mac_addr,omitempty" json:"mac_addr,omitempty"`
	NasPort           int64     `bson:"nas_port,omitempty" json:"nas_port,omitempty,string"`
	NasClass          string    `bson:"nas_class,omitempty" json:"nas_class,omitempty"`
	NasPortId         string    `bson:"nas_port_id,omitempty" json:"nas_port_id,omitempty"`
	NasPortType       int       `bson:"nas_port_type,omitempty" json:"nas_port_type,omitempty"`
	ServiceType       int       `bson:"service_type,omitempty" json:"service_type,omitempty"`
	AcctSessionId     string    `bson:"acct_session_id,omitempty" json:"acct_session_id,omitempty"`
	AcctSessionTime   int       `bson:"acct_session_time,omitempty" json:"acct_session_time,omitempty"`
	AcctInputTotal    int64     `bson:"acct_input_total,omitempty" json:"acct_input_total,omitempty,string"`
	AcctOutputTotal   int64     `bson:"acct_output_total,omitempty" json:"acct_output_total,omitempty,string"`
	AcctInputPackets  int       `bson:"acct_input_packets,omitempty" json:"acct_input_packets,omitempty"`
	AcctOutputPackets int       `bson:"acct_output_packets,omitempty" json:"acct_output_packets,omitempty"`
	AcctStartTime     time.Time `bson:"acct_start_time,omitempty" json:"acct_start_time,omitempty"`
	LastUpdate        time.Time `bson:"last_update,omitempty" json:"last_update,omitempty"`
	AcctStopTime      time.Time `bson:"acct_stop_time,omitempty" json:"acct_stop_time,omitempty"`
}

type Authlog struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string    `bson:"username,omitempty" json:"username,omitempty"`
	NasAddr   string    `bson:"nas_addr,omitempty" json:"nas_addr,omitempty"`
	Cast      int       `bson:"cast,omitempty" json:"cast,omitempty"`
	Result    string    `bson:"result,omitempty" json:"result,omitempty"`
	Reason    string    `bson:"reason,omitempty" json:"reason,omitempty"`
	Timestamp time.Time `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
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
