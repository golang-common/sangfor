// 数据模式

package atrust

// OnlineUserSearch - 在线用户查询参数
// PageSize - 每页的大小 默认20
// PageIndex - 当前页，默认1，从1开始
// Filter - 过滤条件，可选值：all（所有条件）、name（用户名）、displayName（显示名）
// userDirectoryName（所属用户目录）、groupPath（所属组织架构）、remoteIp（接入IP）
// vip（虚拟IP）、os（操作系统类型）、browser（接入方式）
// SearchValue - 过滤搜索值
// SortBy - 排序字段，可选值：name（用户名）、groupPath（所属组织架构）、browser（接入方式）
// os（操作系统类型）、 lastLoginTime（最后登录时间）、authTypeList（认证方式）、 remoteIp（接入IP）
// userDirectoryName（所属用户目录）、vip（虚拟IP）
// Asc - 是否升序：1为升序、0为降序(不传默认为降序)
type OnlineUserSearch struct {
	PageSize    int    `json:"pageSize,omitempty" url:"pageSize,omitempty"`
	PageIndex   int    `json:"pageIndex,omitempty" url:"pageIndex,omitempty"`
	Filter      string `json:"filter,omitempty" url:"filter,omitempty"`
	SearchValue string `json:"searchValue,omitempty" url:"searchValue,omitempty"`
	SortBy      string `json:"sortBy,omitempty" url:"sortBy,omitempty"`
	Asc         int    `json:"asc,omitempty" url:"asc,omitempty"`
}

// OnlineUser 用户在线信息
// Id - 记录ID
// Name - 用户名称
// DisplayName - 显示名
// UserDirectoryName - 用户目录名
// GroupId - 所属组织架构ID
// GroupPath - 所属组织架构路径
// Os - 操作系统信息
// Browser - 浏览器信息
// RemoteIp - 接入IP
// AuthTypeList - 认证类型列表
// LastLoginTime - 最后登录时间
// UserId - 用户ID
// Domain - 登录域
// IsTrusted - TODO: 待确认
// Vips - 虚IP列表
type OnlineUser struct {
	Id                string      `json:"id,omitempty"`
	Name              string      `json:"name,omitempty"`
	DisplayName       string      `json:"displayName,omitempty"`
	UserDirectoryName string      `json:"userDirectoryName,omitempty"`
	GroupId           int         `json:"groupId,omitempty"`
	GroupPath         string      `json:"groupPath,omitempty"`
	Os                string      `json:"os,omitempty"`
	Browser           string      `json:"browser,omitempty"`
	RemoteIp          string      `json:"remoteIp,omitempty"`
	AuthTypeList      []AuthType  `json:"authTypeList,omitempty"`
	LastLoginTime     string      `json:"lastLoginTime,omitempty"`
	UserId            string      `json:"userId,omitempty"`
	Domain            string      `json:"domain,omitempty"`
	IsTrusted         int         `json:"isTrusted,omitempty"`
	Vips              []VirtualIP `json:"vips,omitempty"`
}

// VirtualIP 客户端的虚IP信息
// Ip - 虚IP
// ProxyClientId - 节点ID
// ProxyClientName - 节点名称
type VirtualIP struct {
	Ip              string `json:"ip,omitempty"`
	ProxyClientId   string `json:"proxyClientId,omitempty"`
	ProxyClientName string `json:"proxyClientName,omitempty"`
}

// AuthType 认证类型
// AType - 类型名称
// SubType - 子类型
type AuthType struct {
	AType   string `json:"authType"`
	SubType string `json:"subType"`
}
