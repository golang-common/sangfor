// 数据结构

package ac

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// InsideLib 内置库结构体(病毒库,URL库等)
type InsideLib struct {
	Name      string `json:"name"`       // 库名称
	Type      string `json:"type"`       // 库类型(kav=病毒库,url=URL库,up=网关补丁,contchk=应用识别,trace=审计规则库)
	Current   string `json:"current"`    // 当前版本
	New       string `json:"new"`        // 最新版本
	Expire    string `json:"expire"`     // 升级服务序列号过期时间
	Enable    bool   `json:"enable"`     // 是否启用自动升级
	IsExpired int    `json:"is_expired"` // 规则库是否过期(0=未过期,1=过期)
}

// LogCount 日志统计
type LogCount struct {
	Block  int `json:"block"`  // 拦截日志计数
	Record int `json:"record"` // 记录日志计数
}

// ThroughputFilter 上下行流速请求过滤
type ThroughputFilter struct {
	Unit      string `json:"unit,omitempty"`      // 流量单位(取值bits/bytes)
	Interface string `json:"interface,omitempty"` // 接口名称
}

// Throughput 设备的上行和下行流速返回结构体
type Throughput struct {
	Recv int    `json:"recv"` // 接收流量
	Send int    `json:"send"` // 发出流量
	Unit string `json:"unit"` // 流量单位(bits或bytes)
}

// UserRankFilter 用户流量排行过滤参数
// 过滤字段"groups","users","ips",同时只能选择其中一种过滤条件来过滤,
// 若同时传入多个过滤条件,则过滤条件只会生效1种,优先级为"groups > users > ips"
type UserRankFilter struct {
	Top    int      `json:"top,omitempty"`    // TopN排行
	Line   string   `json:"line,omitempty"`   // 线路号(0:所有线路，1-N:具体线路)
	Groups []string `json:"groups,omitempty"` // 要过滤的组(以"/"开头)
	Users  []string `json:"users,omitempty"`  // 要过滤的用户
	Ips    []string `json:"ips,omitempty"`    // 要过滤的IP(只支持单个IP,不支持IP组)
}

// UserRank 用户流量排行数据
type UserRank struct {
	Id      int              `json:"id"`               // 序号
	Name    string           `json:"name"`             // 用户名
	Group   string           `json:"group"`            // 组
	Ip      string           `json:"ip"`               // IP
	Up      int              `json:"up"`               // 上行流量(bytes)
	Down    int              `json:"down"`             // 下行流量(bytes)
	Total   int              `json:"total"`            // 总流量(bytes)
	Session int              `json:"session"`          // 会话数
	Status  bool             `json:"status"`           // 冻结状态(false为冻结)
	Detail  []UserRankDetail `json:"detail,omitempty"` // 明细
}

func (u *UserRank) UnmarshalJSON(b []byte) error {
	var r = make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	for k, v := range r {
		switch k {
		case "id":
			err = json.Unmarshal(v, &u.Id)
			if err != nil {
				return err
			}
		case "name":
			err = json.Unmarshal(v, &u.Name)
			if err != nil {
				return err
			}
		case "group":
			err = json.Unmarshal(v, &u.Group)
			if err != nil {
				return err
			}
		case "ip":
			err = json.Unmarshal(v, &u.Ip)
			if err != nil {
				return err
			}
		case "up":
			err = json.Unmarshal(v, &u.Up)
			if err != nil {
				return err
			}
		case "down":
			err = json.Unmarshal(v, &u.Down)
			if err != nil {
				return err
			}
		case "total":
			err = json.Unmarshal(v, &u.Total)
			if err != nil {
				return err
			}
		case "session":
			err = json.Unmarshal(v, &u.Session)
			if err != nil {
				return err
			}
		case "status":
			err = json.Unmarshal(v, &u.Status)
			if err != nil {
				return err
			}
		case "detail":
			var a struct {
				Data []UserRankDetail `json:"data"`
			}
			err = json.Unmarshal(v, &a)
			if err != nil {
				return err
			}
			u.Detail = a.Data
		}
	}
	return nil
}

// UserRankDetail 用户流量排行明细
// Id - 用户流量明细中的排序
// App - 应用名称
// Up - 上行字节数
// Down - 下行字节数
// Total - 总字节数
// Percent - 该应用占用户总流量的百分比
type UserRankDetail struct {
	Id      int    `json:"id"`
	App     string `json:"app"`
	Up      int    `json:"up"`
	Down    int    `json:"down"`
	Total   int    `json:"total"`
	Percent int    `json:"percent"`
}

// AppRankFilter 应用流量排行过滤参数
// Top - TopN排行
// Line - 线路号(0:所有线路，1-N:具体线路)
// Groups - 要过滤的组(以"/"开头)
type AppRankFilter struct {
	Top    int      `json:"top,omitempty"`
	Line   string   `json:"line,omitempty"`
	Groups []string `json:"groups,omitempty"`
}

// AppRank 应用流量排行数据
type AppRank struct {
	App       string        `json:"app"`
	Line      int           `json:"line"`
	LineName  string        `json:"line_name"`
	Up        int           `json:"up"`
	Down      int           `json:"down"`
	Total     int           `json:"total"`
	Rate      int           `json:"rate"`
	Session   int           `json:"session"`
	Users     []AppRankUser `json:"users"`
	UserCount int           `json:"user_count"`
}

// AppRankUser 应用流量排行中的用户信息
type AppRankUser struct {
	User  string `json:"user"`
	Grp   string `json:"grp"`
	Ip    string `json:"ip"`
	Up    int    `json:"up"`
	Down  int    `json:"down"`
	Total int    `json:"total"`
}

func (a *AppRank) UnmarshalJSON(b []byte) error {
	var r = make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	for k, v := range r {
		switch k {
		case "app":
			err = json.Unmarshal(v, &a.App)
			if err != nil {
				return err
			}
		case "line":
			err = json.Unmarshal(v, &a.Line)
			if err != nil {
				return err
			}
		case "line_name":
			err = json.Unmarshal(v, &a.LineName)
			if err != nil {
				return err
			}
		case "up":
			err = json.Unmarshal(v, &a.Up)
			if err != nil {
				return err
			}
		case "down":
			err = json.Unmarshal(v, &a.Down)
			if err != nil {
				return err
			}
		case "total":
			err = json.Unmarshal(v, &a.Total)
			if err != nil {
				return err
			}
		case "rate":
			err = json.Unmarshal(v, &a.Rate)
			if err != nil {
				return err
			}
		case "session":
			err = json.Unmarshal(v, &a.Session)
			if err != nil {
				return err
			}
		case "user_data":
			var ud struct {
				Data  []AppRankUser `json:"data,omitempty"`
				Count int           `json:"count,omitempty"`
			}
			err = json.Unmarshal(v, &ud)
			if err != nil {
				return err
			}
			a.UserCount = ud.Count
			a.Users = ud.Data
		}
	}
	return nil
}

// UserBind 用户/IP/MAC绑定关系
// Name - 用户名
// Enable - 是否启用
// Desc - 描述
// AddrType - 绑定类型, ip / mac / ipmac
// Addr - 绑定对象
// - 当 AddrType 为ip时,取值为ip地址(e.g.:192.168.1.1)
// - 当 AddrType 为mac时,取值为mac地址(e.g.:ff-ff-ff-ff-ff-ff)
// - 当 AddrType 为ipmac时,取值为ipmac地址(e.g.:192.168.1.1+ff-ff-ff-ff-ff-ff)
// Limitlogon - 是否启用限制登录
// Noauth - 免认证配置
type UserBind struct {
	Name       string  `json:"name"`
	Enable     bool    `json:"enable,omitempty"`
	Desc       string  `json:"desc,omitempty"`
	AddrType   string  `json:"addr_type,omitempty"`
	Addr       string  `json:"addr,omitempty"`
	Time       Date    `json:"time,omitempty"`
	Limitlogon bool    `json:"limitlogon,omitempty"` // 是否限制登录
	Noauth     *Noauth `json:"noauth,omitempty"`
}

// Noauth 免认证配置
// Enable - 是否启用免认证
// ExpireTime - 免认证过期时间, Unix时间戳, 为0表示永不过期, >0表示过期时 间戳
type Noauth struct {
	Enable     bool `json:"enable,omitempty"`
	ExpireTime int  `json:"expire_time,omitempty"`
}

// BindIpMac IP/MAC绑定关系
// Ip - IP地址
// Mac - MAC地址
// Desc - 描述
type BindIpMac struct {
	Ip   string `json:"ip,omitempty"`
	Mac  string `json:"mac,omitempty"`
	Desc string `json:"desc,omitempty"`
}

// UserAdd 用户增加
type UserAdd struct {
	UserBasic[Time]
	UserExtend[[]UserBindCfg, []string]
}

// UserModify 用户修改
// Name - 用户名(必填)
type UserModify struct {
	Name       string `json:"name"`
	Desc       string `json:"desc,omitempty"`
	ExpireTime Time   `json:"expire_time,omitempty"`
	UserExtend[[]UserBindCfg, *UserLimit]
}

func (u UserModify) MarshalJSON() ([]byte, error) {
	var dmap = make(map[string]any)
	var data = make(map[string]any)
	dmap["name"] = u.Name
	dbytes, err := json.Marshal(u.UserExtend)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dbytes, &data)
	if err != nil {
		return nil, err
	}
	if u.Desc != "" {
		data["desc"] = u.Desc
	}
	if !u.ExpireTime.IsZero() {
		data["expire_time"] = u.ExpireTime
	}
	if len(data) > 0 {
		dmap["data"] = data
	}
	return json.Marshal(dmap)
}

// UserDetail 用户查询，搜索
type UserDetail struct {
	UserBasic[*UserExpire]
	UserExtend[[]map[string]string, *UserLimit]
	Create     string          `json:"create"`
	CreateFlag bool            `json:"create_flag"`
	Policy     []NetPolicyInfo `json:"policy,omitempty"`
}

// UserBasic 用户基本信息
// Name - 用户名
// Desc - 用户描述
// FatherPath - 用户所在组
// ExpireTime - 过期时间
type UserBasic[T interface{ *UserExpire | Time }] struct {
	Name       string `json:"name"`
	Desc       string `json:"desc,omitempty"`
	ShowName   string `json:"show_name,omitempty"`
	FatherPath string `json:"father_path,omitempty"`
	ExpireTime T      `json:"expire_time,omitempty"`
}

// UserExtend 用户扩展信息
// Enable - 是否启用用户
// Desc - 用户描述
// SelfPass - 本地账号密码配置
// BindCfg - 用户绑定信息列表
// CustomCfg - 用户自定义属性键值对
// Logout - 密码认证成功后是否弹出注销窗口
// CommonUser - 用户通用属性
// LimitIpmac - 限制登录地址(IP或MAC,IP支持单个或IP段 (192.168.1.1-192.168.1.2),MAC格式ee-ee-ee-ee-ee-ee)
// TODO: 增加或修改用户时所有的扩展属性都不生效，需要与厂商确认
type UserExtend[
	Bind []UserBindCfg | []map[string]string,
	Limit interface{ *UserLimit | []string },
] struct {
	Enable     bool                `json:"enable"`
	SelfPass   *UserSelfPass       `json:"self_pass,omitempty"`
	BindCfg    Bind                `json:"bind_cfg,omitempty"`
	CustomCfg  []map[string]string `json:"custom_cfg,omitempty"`
	Logout     bool                `json:"logout,omitempty"`
	CommonUser *UserCommon         `json:"common_user,omitempty"`
	LimitIpmac Limit               `json:"limit_ipmac,omitempty"`
}

// UserExpire 用户过期信息
// Enable - 是否启用用户过期
// Date - 过期日期(YYYY-MM-DD)
type UserExpire struct {
	Enable bool   `json:"enable"`
	Date   string `json:"date,omitempty"`
}

// UserLimit 用户登录限制
// Enable - 是否开启限制
// Ipmac - 限制的IP或MAC(IP或MAC,IP支持单个或IP段 (192.168.1.1-192.168.1.2),MAC格式ee-ee-ee-ee-ee-ee)
type UserLimit struct {
	Enable bool     `json:"enable"`
	Ipmac  []string `json:"ipmac"`
}

// UserSelfPass 用户本地密码
// Enable - 是否启用密码
// Password - 密码文本(只在插入时使用，无法从设备读取到)
// ModifyOnce - 初次认证是否修改密码
type UserSelfPass struct {
	Enable     bool   `json:"enable"`
	Password   string `json:"password"`
	ModifyOnce bool   `json:"modify_once"`
}

// UserBindCfg 用户绑定信息
// Ip - e.g:192.168.1.2
// Mac - e.g:ac-ed-ee-ee-ee-ee
// OutTime - 过期时间(e.g:2019-10-31)
// Bindgoal - 绑定方式(noauth:免认证,loginlimit:限制登录,noauth_and_loginlimit:免认证且限制登录)
// Desc - 绑定描述
type UserBindCfg struct {
	Ip       string `json:"ip"`       // e.g:192.168.1.2
	Mac      string `json:"mac"`      // e.g:ac-ed-ee-ee-ee-ee
	OutTime  Date   `json:"out_time"` // 过期时间(e.g:2019-10-31)
	Bindgoal string `json:"bindgoal"` // 绑定方式(noauth:免认证,loginlimit:限制登录,noauth_and_loginlimit:免认证且限制登录)
	Desc     string `json:"desc"`     // 绑定描述
}

// UserCommon 用户通用属性
// AllowChange - 是否允许修改本地密码
// Enable - 是否允许多人使用该账号登录
type UserCommon struct {
	AllowChange bool `json:"allow_change"`
	Enable      bool `json:"enable"`
}

// UserSearch 用户搜索传入结构体(最多返回100个)
// SearchType - 搜索类型(user/ip/mac)
// SearchValue - 搜索值
// - 当类型为user时,搜索值为用户名(支持模糊搜索),e.g.:张三
// - 当类型为ip时,搜索用户IP段,e.g.:{"start":"1.1.1.1","end":"1.1.1.10"}
// - 当类型为mac时,搜索用户绑定mac地址,e.g.:ee-ee-ee-ee-ee-ee
// Extend - 搜索扩展字段
type UserSearch struct {
	SearchType  string        `json:"search_type"`
	SearchValue interface{}   `json:"search_value"`
	Extend      *UserSearchEX `json:"extend,omitempty"`
}

// UserSearchEX 搜索扩展字段
// FatherPath - 指定搜索father_path组中的用户,默认为"/"
// CustomCfg - 自定义属性的键值对(不支持同时搜索多个自定义属性)
// UserStatus - 用户状态(共有3种,all:启用和禁用 enabled:启用 disabled:禁用,默认为"all")
// Public - true:搜索过滤出允许多人同时使用的帐号,默认为false
// Expire - 账号过期时间(start:起始时间 end:结束时间 start和end成 对出现,组成时间段)
type UserSearchEX struct {
	FatherPath string            `json:"father_path,omitempty"`
	CustomCfg  map[string]string `json:"custom_cfg,omitempty"`
	UserStatus string            `json:"user_status,omitempty"`
	Public     bool              `json:"public,omitempty"`
	Expire     *TimeRange        `json:"expire,omitempty"`
}

// UserPolicyOper 用户策略操作
// User - 需要修改策略的用户
// Opr - 操作字段
// - add:在原有策略上增加
// - del:在原有策略上删除(无法删除 应用所有用户的策略，无法删除只应用单个用户的策略)
// - modify:将策略设置为policy字段所指定的,会清除原有策略
// Dn - 域名
// Policy - 策略名列表
type UserPolicyOper struct {
	User   string   `json:"user,omitempty"`
	Opr    string   `json:"opr,omitempty"`
	Dn     string   `json:"dn,omitempty"`
	Policy []string `json:"policy,omitempty"`
}

// GroupPolicyOper 用户策略操作
// Group - 需要修改策略的组
// Opr - 操作字段
// - add:在原有策略上增加
// - del:在原有策略上删除(无法删除 应用所有用户的策略，无法删除只应用单个用户的策略)
// - modify:将策略设置为policy字段所指定的,会清除原有策略
// Dn - 域名
// Policy - 策略名列表
type GroupPolicyOper struct {
	Group  string   `json:"group,omitempty"`
	Opr    string   `json:"opr,omitempty"`
	Dn     string   `json:"dn,omitempty"`
	Policy []string `json:"policy,omitempty"`
}

// Group 组信息
// Path - 组路径,最多支持15层级目录创建(以"/"开头,且不支持向域 用户组添加组)
// Desc - 组描述
type Group struct {
	Path string `json:"path"`
	Desc string `json:"desc,omitempty"`
}

type TimeRange struct {
	Start Time `json:"start,omitempty"`
	End   Time `json:"end,omitempty"`
}

type IpRange struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

// NetPolicy 上网策略结构
// PolicyInfo - 策略信息
// UserInfo - 策略关联的用户信息
type NetPolicy struct {
	PolicyInfo NetPolicyInfo     `json:"policy_info,omitempty"`
	UserInfo   NetPolicyUserInfo `json:"user_info,omitempty"`
}

// NetPolicyInfo 上网策略信息
// Name - 策略名
// Type - 策略类型
// Founder - 策略创建者
// Expire - 过期时间
// Status - 是否启用
// Depict - 策略描述信息
type NetPolicyInfo struct {
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Founder string `json:"founder,omitempty"`
	Expire  string `json:"expire,omitempty"`
	Status  bool   `json:"status,omitempty"`
	Depict  string `json:"depict,omitempty"`
}

// NetPolicyUserInfo 上网策略关联的用户信息
// Ou          - 在线用户信息
// Aduser      - 域用户信息
// Adgroup     - 域安全组信息
// ExcAduser   - 排除域用户信息
// Attribute   - 域属性信息
// UserAttrGrp - 用户,组属性信息
// Sourceip    - 源IP
// Location    - 位置列表
// Terminal    - 终端列表
// TargetArea  - 目标区域
// Local       - 关联(适用)的用户
type NetPolicyUserInfo struct {
	Ou          []string `json:"ou,omitempty"`
	Aduser      []string `json:"aduser,omitempty"`
	Adgroup     []string `json:"adgroup,omitempty"`
	ExcAduser   []string `json:"exc_aduser,omitempty"`
	Attribute   []string `json:"attribute,omitempty"`
	UserAttrGrp []string `json:"user_attr_grp,omitempty"`
	Sourceip    []string `json:"sourceip,omitempty"`
	Location    []string `json:"location,omitempty"`
	Terminal    []string `json:"terminal,omitempty"`
	TargetArea  []string `json:"target_area,omitempty"`
	Local       string   `json:"local,omitempty"`
}

// FluxPolicy 流控策略结构
// Id - 通道ID
// Name - 通道名
// FatherId - 父通道名称
// IpGroup - 目标IP组(多个逗号分隔)
// Object - 适用对象(多个逗号分隔,位置/用户/终端...)
// Service - 适用应用(多个逗号分隔)
// Time - 生效时间(e.g.:全天)
// Status - 策略是否启用,true:启用,false:禁用
// Assured - 保证带宽，数组包含上行和下行，-1表示无限制
// Max - 最大带宽，数组包含上行和下行，-1表示无限制
// Single - 单用户限制带宽，数组包含上行和下行，-1表示无限制
// IsDefaultChild - 是否为默认通道,true:是默认通道,false:不是默认通道
// Childrens - 子通道对象数组
// IsLowSpeed - 域属性信息
// TargetUsers - 目标用户(多个用户逗号分隔)
// IpGroup - 目标IP组(多组逗号分隔)
// FIXME: 注释中的字段均为API文档中有但实际请求中不一致或不存在的值,需厂商更新API(深信服垃圾)
type FluxPolicy struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	FatherId string   `json:"fatherID,omitempty"`
	IpGroup  string   `json:"di,omitempty"`
	Object   string   `json:"object,omitempty"`
	Service  string   `json:"service,omitempty"`
	Time     string   `json:"time,omitempty"`
	Status   bool     `json:"status,omitempty"`
	Assured  []string `json:"assured,omitempty"`
	Max      []string `json:"max,omitempty"`
	Single   []string `json:"single,omitempty"`
	//IsDefaultChild bool          `json:"is_default_child,omitempty"`
	//Childrens      []*FluxPolicy `json:"childrens,omitempty"`
	//IsLowSpeed     []string      `json:"is_low_speed,omitempty"`
	//TargetUsers    string        `json:"target_users,omitempty"`
	//IpGroup        string        `json:"ip_group,omitempty"`
}

// OnlineFilter 获取在线用户是的过滤选项
// Status - 用户状态(all-所有,frozen-已冻结,active-活跃)
// Terminal - all-所有 pc-PC用户 mobile-移动终端 multi-多终端 iot-哑终端 armarium-医疗设备 custom-用户自定义设备
// FilterType - 搜索类型(user-用户组名,ip-IP地址数组,mac-mac地址数组)
// FilterValue - 与搜索类型对应的值数组(用户名支持模糊查询)
type OnlineFilter struct {
	Status      string   `json:"status,omitempty"`
	Terminal    string   `json:"terminal,omitempty"`
	FilterType  string   `json:"type,omitempty"`
	FilterValue []string `json:"value,omitempty"`
}

func (o OnlineFilter) MarshalJSON() ([]byte, error) {
	var r = make(map[string]any)
	if o.Status != "" {
		r["status"] = o.Status
	}
	if o.Terminal != "" {
		r["terminal"] = o.Terminal
	}
	if o.FilterType != "" && len(o.FilterValue) > 0 {
		r["filter"] = map[string]any{
			"type":  o.FilterType,
			"value": o.FilterValue,
		}
	}
	return json.Marshal(r)
}

// OnlineUser 在线用户结构
// Name - 用户名
// ShowName - 显示名
// FatherPath - 组路径(在外发上线请求时，字段名会被解析为group)
// Ip - 用户当前IP地址
// Mac - 用户当前MAC地址
// Terminal - 与 OnlineFilter.Terminal 中定义一致
// Authway - 认证方式，0为不需要认证，1为密码认证，2为短信认证，3为单点 登录，4为免认证
// LoginTime - 登录时间戳(Unix 时间戳)
// OnlineTime - 在线时长，单位秒
type OnlineUser struct {
	Name       string `json:"name,omitempty"`
	ShowName   string `json:"show_name,omitempty"`
	FatherPath string `json:"father_path,omitempty"`
	Ip         string `json:"ip,omitempty"`
	Mac        string `json:"mac,omitempty"`
	Terminal   int    `json:"terminal,omitempty"`
	Authway    int    `json:"authway,omitempty"`
	LoginTime  int    `json:"login_time,omitempty"`
	OnlineTime int    `json:"online_time,omitempty"`
}

func (d OnlineUser) MarshalJSON() ([]byte, error) {
	var r = make(map[string]any)
	v := reflect.ValueOf(d)
	for i := 0; i < v.NumField(); i++ {
		tg := v.Type().Field(i).Tag
		j := tg.Get("json")
		jl := strings.Split(j, ",")
		var a = jl[0]
		if a == "father_path" {
			a = "group"
		}
		if !v.Field(i).IsZero() {
			r[a] = v.Field(i).Interface()
		}
	}
	return json.Marshal(r)
}

// IpMacBind IP/MAC绑定关系
type IpMacBind struct {
	Ip   string `json:"ip,omitempty"`
	Mac  string `json:"mac,omitempty"`
	Desc string `json:"desc,omitempty"`
}

type Date struct {
	time.Time
}

func (d Date) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(d.Format(`"2006-01-02"`)), nil
}

func (d Date) String() string {
	return d.Format(`2006-01-02`)
}

// Time 符合AC要求的时间格式
type Time struct {
	time.Time
}

func (d Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format(`2006-01-02 15:04:05`))), nil
}

func (d Time) String() string {
	return d.Format(`2006-01-02 15:04:05`)
}
