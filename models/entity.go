package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Config struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id" query:"id" form:"id"`
	Type   string             `bson:"type" json:"type" form:"type" query:"type"`
	Name   string             `bson:"name" json:"name" form:"name" query:"name"`
	Value  string             `bson:"value" json:"value" form:"value" query:"value"`
	Remark string             `bson:"remark" json:"remark" form:"remark" query:"remark"`
}

type Operator struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id" query:"id" form:"id"`
	Email     string             `bson:"email,omitempty" json:"email" form:"email" query:"email"`
	Username  string             `bson:"username,omitempty" json:"username" form:"username" query:"username"`
	Password  string             `bson:"password,omitempty" json:"password" form:"password" query:"password"`
	ApiSecret string             `bson:"api_secret,omitempty" json:"api_secret" form:"api_secret" query:"api_secret"`
	MfaSecret string             `bson:"mfa_secret,omitempty" json:"mfa_secret" form:"mfa_secret" query:"mfa_secret"`
	Perms     []string           `bson:"perms,omitempty" json:"perms" form:"perms" query:"perms"`
	Status    string             `bson:"status,omitempty" json:"status" form:"status" query:"status"`
	Remark    string             `bson:"remark,omitempty" json:"remark" form:"remark" query:"remark"`
}

type OprLog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id" query:"id" form:"id"`
	Oprname   string             `bson:"oprname,omitempty" json:"oprname" form:"oprname" query:"oprname"`
	Srcip     string             `bson:"srcip,omitempty" json:"srcip" form:"srcip" query:"srcip"`
	Action    string             `bson:"action,omitempty" json:"action" form:"action" query:"action"`
	Remark    string             `bson:"remark,omitempty" json:"remark" form:"remark" query:"remark"`
	Timestamp primitive.DateTime `bson:"timestamp,omitempty" json:"timestamp" form:"-" query:"-"`
}

type Vpe struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id" query:"id" form:"id"`
	Identifier string             `bson:"identifier,omitempty" json:"identifier" form:"identifier" query:"identifier"`
	Name       string             `bson:"name,omitempty" json:"name" form:"name" query:"name"`
	Ipaddr     string             `bson:"ipaddr,omitempty" json:"ipaddr" form:"ipaddr" query:"ipaddr"`
	Secret     string             `bson:"secret,omitempty" json:"-" form:"secret" query:"secret"`
	VendorCode string             `bson:"vendor_code,omitempty" json:"vendor_code" form:"vendor_code" query:"vendor_code"`
	CoaPort    int                `bson:"coa_port,omitempty" json:"coa_port" form:"coa_port" query:"coa_port"`
	Status     string             `bson:"status,omitempty" json:"status" form:"status" query:"status"`
	LdapId     primitive.ObjectID `bson:"ldap_id,omitempty" json:"ldap_id,string" form:"ldap_id" query:"ldap_id"`
	Remark     string             `bson:"remark,omitempty" json:"remark" form:"remark" query:"remark"`
}

type Ldap struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id" query:"id" form:"id"`
	Name       string             `bson:"name,omitempty" json:"name" form:"name" query:"name"`
	Address    string             `bson:"address,omitempty" json:"address" form:"address" query:"address"`
	Password   string             `bson:"password,omitempty" json:"-" form:"password" query:"password"`
	Searchdn   string             `bson:"searchdn,omitempty" json:"searchdn" form:"searchdn" query:"searchdn"`
	Basedn     string             `bson:"basedn,omitempty" json:"basedn" form:"basedn" query:"basedn"`
	UserFilter string             `bson:"user_filter,omitempty" json:"user_filter" form:"user_filter" query:"user_filter"`
	Istls      string             `bson:"istls,omitempty" json:"istls" form:"istls" query:"istls"`
	Status     string             `bson:"status,omitempty" json:"status" form:"status" query:"status"`
	Remark     string             `bson:"remark,omitempty" json:"remark" form:"remark" query:"remark"`
}

type Profile struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id" query:"id" form:"id"`
	Name            string             `bson:"name" json:"name" form:"name" query:"name"`
	Domain          string             `bson:"domain" json:"domain" form:"domain" query:"domain"`
	InterimInterval int                `bson:"interim_interval" json:"interim_interval" form:"interim_interval" query:"interim_interval"`
	AddrPool        string             `bson:"addr_pool" json:"addr_pool" form:"addr_pool" query:"addr_pool"`
	ActiveNum       int                `bson:"active_num" json:"active_num" form:"active_num" query:"active_num"`
	MfaStatus       string             `bson:"mfa_status" json:"mfa_status" form:"mfa_status" query:"mfa_status"`
	UpRate          int                `bson:"up_rate" json:"up_rate" form:"up_rate" query:"up_rate"`
	DownRate        int                `bson:"down_rate" json:"down_rate" form:"down_rate" query:"down_rate"`
	LimitPolicy     string             `bson:"limit_policy" json:"limit_policy" form:"limit_policy" query:"limit_policy"`
	UpLimitPolicy   string             `bson:"up_limit_policy" json:"up_limit_policy" form:"up_limit_policy" query:"up_limit_policy"`
	DownLimitPolicy string             `bson:"down_limit_policy" json:"down_limit_policy" form:"down_limit_policy" query:"down_limit_policy"`
	BillTimes       int                `bson:"bill_times" json:"bill_times" form:"bill_times" query:"bill_times"`
	BillTimeunit    string             `bson:"bill_timeunit" json:"bill_timeunit" form:"bill_timeunit" query:"bill_timeunit"`
	Status          string             `bson:"status" json:"status" form:"status" query:"status"`
	Remark          string             `bson:"remark" json:"remark" form:"remark" query:"remark"`
}

type Subscribe struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty" json:"id" query:"id" form:"id"`
	VpeSids    []primitive.ObjectID `bson:"vpe_sids,omitempty" json:"vpe_sids" form:"vpe_sids" query:"vpe_sids"`
	Profile    Profile              `bson:"profile,omitempty" json:"profile,string" form:"profile" query:"profile"`
	Realname   string               `bson:"realname,omitempty" json:"realname" form:"realname" query:"realname"`
	Email      string               `bson:"email,omitempty" json:"email" form:"email" query:"email"`
	Username   string               `bson:"username,omitempty" json:"username" form:"username" query:"username"`
	Password   string               `bson:"password,omitempty" json:"-" form:"password" query:"password"`
	MfaStatus  string               `bson:"mfa_status,omitempty" json:"mfa_status" form:"mfa_status" query:"mfa_status"`
	MfaSecret  string               `bson:"mfa_secret,omitempty" json:"mfa_secret" form:"mfa_secret" query:"mfa_secret"`
	Ipaddr     string               `bson:"ipaddr,omitempty" json:"ipaddr" form:"ipaddr" query:"ipaddr"`
	Macaddr    string               `bson:"macaddr,omitempty" json:"macaddr" form:"macaddr" query:"macaddr"`
	Vlanid1    int                  `bson:"vlanid_1,omitempty" json:"vlanid1" form:"vlanid1" query:"vlanid1"`
	Vlanid2    int                  `bson:"vlanid_2,omitempty" json:"vlanid2" form:"vlanid2" query:"vlanid2"`
	BindMac    int                  `bson:"bind_mac,omitempty" json:"bind_mac" form:"bind_mac" query:"bind_mac"`
	BindVlan   int                  `bson:"bind_vlan,omitempty" json:"bind_vlan" form:"bind_vlan" query:"bind_vlan"`
	Status     string               `bson:"status,omitempty" json:"status" form:"status" query:"status"`
	Remark     string               `bson:"remark,omitempty" json:"remark" form:"remark" query:"remark"`
	ExpireTime primitive.DateTime   `bson:"expire_time,omitempty" json:"expire_time" form:"-" query:"-"`
	Timestamp  primitive.DateTime   `bson:"timestamp,omitempty" json:"timestamp" form:"-" query:"-"`
}

type Accounting struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id" query:"id" form:"id"`
	Username          string             `bson:"username,omitempty" json:"username" form:"username" query:"username"`
	NasId             string             `bson:"nas_id,omitempty" json:"nas_id" form:"nas_id" query:"nas_id"`
	NasAddr           string             `bson:"nas_addr,omitempty" json:"nas_addr" form:"nas_addr" query:"nas_addr"`
	NasPaddr          string             `bson:"nas_paddr,omitempty" json:"nas_paddr" form:"nas_paddr" query:"nas_paddr"`
	SessionTimeout    int                `bson:"session_timeout,omitempty" json:"session_timeout" form:"session_timeout" query:"session_timeout"`
	FramedIpaddr      string             `bson:"framed_ipaddr,omitempty" json:"framed_ipaddr" form:"framed_ipaddr" query:"framed_ipaddr"`
	FramedNetmask     string             `bson:"framed_netmask,omitempty" json:"framed_netmask" form:"framed_netmask" query:"framed_netmask"`
	MacAddr           string             `bson:"mac_addr,omitempty" json:"mac_addr" form:"mac_addr" query:"mac_addr"`
	NasPort           int64              `bson:"nas_port,omitempty" json:"nas_port,string" form:"nas_port" query:"nas_port"`
	NasClass          string             `bson:"nas_class,omitempty" json:"nas_class" form:"nas_class" query:"nas_class"`
	NasPortId         string             `bson:"nas_port_id,omitempty" json:"nas_port_id" form:"nas_port_id" query:"nas_port_id"`
	NasPortType       int                `bson:"nas_port_type,omitempty" json:"nas_port_type" form:"nas_port_type" query:"nas_port_type"`
	ServiceType       int                `bson:"service_type,omitempty" json:"service_type" form:"service_type" query:"service_type"`
	AcctSessionId     string             `bson:"acct_session_id,omitempty" json:"acct_session_id" form:"acct_session_id" query:"acct_session_id"`
	AcctSessionTime   int                `bson:"acct_session_time,omitempty" json:"acct_session_time" form:"acct_session_time" query:"acct_session_time"`
	AcctInputTotal    int64              `bson:"acct_input_total,omitempty" json:"acct_input_total,string" form:"acct_input_total" query:"acct_input_total"`
	AcctOutputTotal   int64              `bson:"acct_output_total,omitempty" json:"acct_output_total,string" form:"acct_output_total" query:"acct_output_total"`
	AcctInputPackets  int                `bson:"acct_input_packets,omitempty" json:"acct_input_packets" form:"acct_input_packets" query:"acct_input_packets"`
	AcctOutputPackets int                `bson:"acct_output_packets,omitempty" json:"acct_output_packets" form:"acct_output_packets" query:"acct_output_packets"`
	AcctStartTime     primitive.DateTime `bson:"acct_start_time,omitempty" json:"acct_start_time" form:"-" query:"-"`
	LastUpdate        primitive.DateTime `bson:"last_update,omitempty" json:"last_update" form:"-" query:"-"`
	AcctStopTime      primitive.DateTime `bson:"acct_stop_time,omitempty" json:"acct_stop_time" form:"-" query:"-"`
}

type Authlog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id" query:"id" form:"id"`
	Username  string             `bson:"username,omitempty" json:"username" form:"username" query:"username"`
	NasAddr   string             `bson:"nas_addr,omitempty" json:"nas_addr" form:"nas_addr" query:"nas_addr"`
	Cast      int                `bson:"cast,omitempty" json:"cast" form:"cast" query:"cast"`
	Result    string             `bson:"result,omitempty" json:"result" form:"result" query:"result"`
	Reason    string             `bson:"reason,omitempty" json:"reason" form:"reason" query:"reason"`
	Timestamp primitive.DateTime `bson:"timestamp,omitempty" json:"timestamp" form:"-" query:"-"`
}



// AuthorizationProfile Method

func (a Subscribe) GetExpireTime() time.Time {
	return a.ExpireTime.Time()
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
