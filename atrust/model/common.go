package model

type SearchRecord struct {
	Creator    string `json:"creator,omitempty"`
	DataType   string `json:"dataType,omitempty"`
	ServerName string `json:"serverName,omitempty"`
	EntityType string `json:"entityType,omitempty"`
	Domain     string `json:"domain,omitempty"`
}

type CommonRecord struct {
	CreatorRole string `json:"creatorRole,omitempty"`
	Creator     string `json:"creator,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

// EditResourceGroup 变更对象时的应用组操作,编辑时ID和名称只要传一个,都传以ID为准
// ResourceGroupEditWay - 编辑方式 "set"=重置, “append”=追加, "delete"=删除
// ResourceGroupIdList - 资源组ID列表
// ResourceGroupNameList - 资源组名列表
type EditResourceGroup struct {
	ResourceGroupEditWay  string   `json:"resourceGroupEditWay,omitempty"`
	ResourceGroupIdList   []string `json:"resourceGroupIdList,omitempty"`
	ResourceGroupNameList []string `json:"resourceGroupNameList,omitempty"`
}

// EditResource 变更对象时的应用操作,编辑时ID和名称只要传一个,都传以ID为准
// ResourceEditWay - 编辑方式 "set"=重置, “append”=追加, "delete"=删除
// ResourceIdList - 应用ID列表
// ResourceNameList - 应用名列表
type EditResource struct {
	ResourceEditWay  string   `json:"resourceEditWay,omitempty"`
	ResourceIdList   []string `json:"resourceIdList,omitempty"`
	ResourceNameList []string `json:"resourceNameList,omitempty"`
}

// EditRole 变更对象时的角色操作,编辑时ID和名称只要传一个,都传以ID为准
// RoleEditWay - 编辑方式 "set"=重置, “append”=追加, "delete"=删除
// RoleIdList - 角色ID列表
// RoleNameList - 角色名列表
type EditRole struct {
	RoleEditWay  string   `json:"roleEditWay,omitempty"`
	RoleIdList   []string `json:"roleIdList,omitempty"`
	RoleNameList []string `json:"roleNameList,omitempty"`
}

// Resource 应用
// AccessModel - 访问模式，WEB/L3VPN/TP-WEB
// SubModel - 子访问模式, WEB/L3VPN/TP-WEB
// DataType - group / resource 应用组/应用
// Status - 启用状态，0禁用，1启用
// GroupId - 所属应用分类ID
// GroupName - 所属应用分类名称
// NodeGroupId - 应用所属节点区域的id, 综合网关无需填写, 分离式设备必填
// NodeGroupName - 应用所属节点区域的名称, 综合网关无需填写, 分离式设备必填
// IconId - 预置的图标ID，"1"~"31"可供选择
// AllowApply - 是否允许用户自助申请此应用, 0=不允许, 1=允许
// AppAddress - 后端应用服务器地址, 示例：https://1.2.3.1
// AliasAppAddress - 后端应用服务器地址别名，英文逗号分隔, 示例： https://1.1.1.2,https://1.1.1.3,https://1.1.1.4
// AccessAddress - 前端应用访问地址,示例：https://testweb.com
// TrustedCertId - 授信证书ID，默认内置web应用证书ID为"default"
// Ext - 扩展属性
type Resource struct {
	Id              string       `json:"id,omitempty"`
	Name            string       `json:"name,omitempty"`
	Description     string       `json:"description,omitempty"`
	DataType        string       `json:"dataType,omitempty"`
	AccessModel     string       `json:"accessModel,omitempty"`
	SubModel        string       `json:"subModel,omitempty"`
	Status          int          `json:"status,omitempty"`
	GroupId         string       `json:"groupId,omitempty"`
	GroupName       string       `json:"groupName,omitempty"`
	NodeGroupId     string       `json:"nodeGroupId,omitempty"`
	NodeGroupName   string       `json:"nodeGroupName,omitempty"`
	IconId          string       `json:"iconId,omitempty"`
	AllowApply      int          `json:"allowApply,omitempty"`
	ApplyForInfo    ResApplyInfo `json:"applyForInfo,omitempty"`
	AppAddress      string       `json:"appAddress,omitempty"`
	AliasAppAddress string       `json:"aliasAppAddress,omitempty"`
	AccessAddress   string       `json:"accessAddress,omitempty"`
	TrustedCertId   string       `json:"trustedCertId,omitempty"`
	Ext             WebResExtend `json:"ext,omitempty"`
}

// WebResExtend Web应用的扩展信息
// Hide - 是否在用户应用中心隐藏，0不隐藏，1隐藏
// PathStatus - url路径规则模式, 0黑名单模式，1白名单模式
// Paths - url路径规则, 例如： ["/asdsad", "/asdasdwe"]
// Security - 应用安全配置
type WebResExtend struct {
	Hide        int      `json:"hide,omitempty"`
	PathStatus  int      `json:"pathStatus"`
	Paths       []string `json:"paths,omitempty"`
	DependSites struct {
		Enable    string `json:"enable,omitempty"`
		PanDomain string `json:"panDomain,omitempty"`
		SiteList  []struct {
			AppAddress    string `json:"appAddress,omitempty"`
			AccessAddress string `json:"accessAddress,omitempty"`
		} `json:"siteList,omitempty"`
	} `json:"dependSites,omitempty"`
	Security struct {
		Watermmark struct {
			Enable bool `json:"enable,omitempty"`
		} `json:"watermmark,omitempty"`
	} `json:"security,omitempty"`
	OpenModel WebResOpenModel `json:"openModel,omitempty"`
}

// WebResOpenModel web应用打开模式
// Model - 方式，有如下几种值：no，default-browser，custom-browser
// ProgramName - 程序名称，如果model填写的custom-browser是，则此项需要填写，有如下几种值：Google Chrome（谷歌浏览器）、Internet Explorer（IE浏览器）、360安全浏览器、360极速浏览器、Firefox（火狐浏览器）、Microsoft Edge、QQ浏览器、搜狗浏览器
// ProcessName - 进程名，如果填写了programName，则需要填写此项，每种浏览器对应的进程名为： Google Chrome（谷歌浏览器）：chrome.exe Internet Explorer（IE浏览器）：iexplore.exe 360安全浏览器：360se.exe 360极速浏览器：360chrome.exe Firefox（火狐浏览器）：firefox.exe  Microsoft Edge：msedge.exe QQ浏览器：QQBrowser.exe 搜狗浏览器：SogouExplorer.exe
// IssuerName - 签名者名称，如果填写了programName，则需要填写此项，每种浏览器对应的签名者名称为： Google Chrome（谷歌浏览器）：Google LLC  Internet Explorer（IE浏览器）：Microsoft Corporation ,360安全浏览器：Beijing Qihu Technology Co., Ltd. ,360极速浏览器：Beijing Qihu Technology Co., Ltd. ,Firefox（火狐浏览器）：Mozilla Corporation ,Microsoft Edge：Microsoft Corporation ,QQ浏览器：Tencent Technology(Shenzhen) Company Limited , 搜狗浏览器：Beijing Sogou Technology Development Co., Ltd.
// OriginalFilename - 原始文件名，如果填写了programName，则需要填写此项，每种浏览器对应的原始文件名为： ,Google Chrome（谷歌浏览器）：chrome.exe ,Internet Explorer（IE浏览器）：IEXPLORE.EXE ,360安全浏览器：360se.exe ,360极速浏览器：360chrome.exe ,Firefox（火狐浏览器）：firefox.exe ,Microsoft Edge：msedge.exe ,QQ浏览器：QQBrowser.exe ,搜狗浏览器：（该浏览器不存在原始文件名）
// UseDefaultBrowser - 允许在未找到该浏览器时自动打开默认浏览器，0不允许，1允许
type WebResOpenModel struct {
	Model             string   `json:"model,omitempty"`
	ProgramName       string   `json:"programName,omitempty"`
	ProcessName       string   `json:"processName,omitempty"`
	IssuerName        []string `json:"issuerName,omitempty"`
	OriginalFilename  string   `json:"originalFilename,omitempty"`
	UseDefaultBrowser int      `json:"useDefaultBrowser,omitempty"`
}

// ResApplyInfo 应用用户申请信息
// Content - 告警内容, 例如："您当前没有该应用的访问权限，请向管理员申请
// AccessReason - 申请理由，最少2条，最多5条, 例如：["我因为岗位职责原因需要访问该应用", "我因为个人特殊原因需要访问该应用"]
type ResApplyInfo struct {
	Content      string   `json:"content,omitempty"`
	AccessReason []string `json:"accessReason,omitempty"`
}

// ResGroup 应用组
// Name - 组名
// ParentGroupId - 父组ID
// Status - 状态, 1=启用
// SequenceNumber - 顺序号, 仅查询
type ResGroup struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	ParentGroupId  string `json:"parentGroupId,omitempty"`
	Status         int    `json:"status,omitempty"`
	SequenceNumber string `json:"sequenceNumber,omitempty"`
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

type UpdateRole struct {
	RoleEditWay  string   `json:"roleEditWay,omitempty"`
	RoleIdList   []string `json:"roleIdList,omitempty"`
	RoleNameList []string `json:"roleNameList,omitempty"`
}

type UpdateResource struct {
	ResourceEditWay  string   `json:"resourceEditWay,omitempty"`
	ResourceIdList   []string `json:"resourceIdList,omitempty"`
	ResourceNameList []string `json:"resourceNameList,omitempty"`
}

type UpdateResourceGroup struct {
	ResourceGroupEditWay  string   `json:"resourceGroupEditWay,omitempty"`
	ResourceGroupIdList   []string `json:"resourceGroupIdList,omitempty"`
	ResourceGroupNameList []string `json:"resourceGroupNameList,omitempty"`
}

type UpdateUser struct {
	UserIdEditWay string   `json:"userIdEditWay,omitempty"`
	UserIdList    []string `json:"userIdList,omitempty"`
	UserNameList  []string `json:"userNameList,omitempty"`
}

type UpdateGroup struct {
	GroupIdEditWay string   `json:"groupIdEditWay,omitempty"`
	GroupIdList    []string `json:"groupIdList,omitempty"`
	GroupNameList  []string `json:"groupNameList,omitempty"`
}

type CommonArg struct {
	Id                string   `json:"id,omitempty" url:"id,omitempty"`
	IdList            []string `json:"idList,omitempty" url:"idList,omitempty"`
	Name              string   `json:"name,omitempty" url:"name,omitempty"`
	NameList          []string `json:"nameList,omitempty" url:"nameList,omitempty"`
	Path              string   `json:"path,omitempty" url:"path,omitempty"`
	Type              int      `json:"type,omitempty" url:"type,omitempty"`
	UserDirectoryId   string   `json:"userDirectoryId,omitempty" url:"userDirectoryId,omitempty"`
	UserDirectoryName string   `json:"userDirectoryName,omitempty" url:"userDirectoryName,omitempty"`
	Value             any      `json:"value,omitempty" url:"-"`
}
