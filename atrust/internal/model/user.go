package model

type User struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	ExpiredTime string `json:"expiredTime,omitempty"`
	Status      int    `json:"status,omitempty"`
	ExternalId  string `json:"externalId,omitempty"`
}

type UserEntity struct {
	User
	SearchRecord
	CommonRecord
	LastLoginTime  string `json:"lastLoginTime,omitempty"`
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	Path           string `json:"path,omitempty"`
	AutoCompose    IdName `json:"authCompose,omitempty"`
	UserPolicy     IdName `json:"userPolicy,omitempty"`
}

type LocalUserAdd struct {
	UserAdd
	GroupId  string `json:"groupId,omitempty"`
	Password string `json:"password,omitempty"`
	PwdModel string `json:"pwdModel,omitempty"`
}

type ExtUserAdd struct {
	UserAdd
	UserDirectoryId   string       `json:"userDirectoryId,omitempty"`
	UserDirectoryName string       `json:"userDirectoryName,omitempty"`
	Ext               UserExternal `json:"ext,omitempty"`
}

type UserAdd struct {
	User
	InheritGroup  int      `json:"inheritGroup,omitempty"`
	AuthComposeId string   `json:"authComposeId,omitempty"`
	UserPolicyId  string   `json:"userPolicyId,omitempty"`
	RoleIdList    []string `json:"roleIdList,omitempty"`
	RoleNameList  []string `json:"roleNameList,omitempty"`
}

type LocalUserUpdate struct {
	UserUpdate
	GroupId  string `json:"groupId,omitempty"`
	PwdModel string `json:"pwdModel,omitempty"`
	Password string `json:"password,omitempty"`
}

type ExtUserUpdate struct {
	UserUpdate
	Ext               UserExternal `json:"ext,omitempty"`
	UserDirectoryId   string       `json:"userDirectoryId,omitempty"`
	UserDirectoryName string       `json:"userDirectoryName,omitempty"`
}

type ExtUserBatchUpdate struct {
	UserUpdate
	Ext UserExternal `json:"ext,omitempty"`
}

type UserUpdate struct {
	InheritGroup    int    `json:"inheritGroup,omitempty"`
	AuthComposeId   string `json:"authComposeld,omitempty"`
	AuthComposeName string `json:"authComposeName,omitempty"`
	UserPolicyId    string `json:"userPolicyId,omitempty"`
	UserPolicyName  string `json:"userPolicyName,omitempty"`
	UpdateRole
	UpdateResource
	UpdateResourceGroup
}

type LocalUserDetail struct {
	UserDetail
	Group         IdName `json:"group,omitempty"`
	AdminRoleName string `json:"adminRoleName,omitempty"`
}

type ExtUserDetail struct {
	UserDetail
	Ext         UserExternal `json:"ext,omitempty"`
	InheritRole int          `json:"inheritRole,omitempty"`
	IsDelete    int          `json:"isDelete,omitempty"`
}

type UserDetail struct {
	User
	CommonRecord
	AuthCompose    IdName     `json:"authComposed,omitempty"`
	UserPolicy     IdName     `json:"userPolicy,omitempty"`
	ResourceList   []Resource `json:"resourceList,omitempty"`
	RoleList       []IdName   `json:"roleList,omitempty"`
	Password       string     `json:"password,omitempty"`
	LastLoginTime  string     `json:"lastLoginTime,omitempty"`
	LastUpdateTime string     `json:"lastUpdateTime,omitempty"`
	InheritGroup   int        `json:"inheritGroup,omitempty"`
}

type UserExternal struct {
	FieldDataSource UserExtFieldDataSource `json:"fieldDataSource,omitempty"`
}

type UserExtFieldDataSource struct {
	Description string `json:"description,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Email       string `json:"email,omitempty"`
	ExpiredTime string `json:"expiredTime,omitempty"`
	Path        string `json:"path,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Status      int    `json:"status,omitempty"`
}

type ResourceRelation struct {
	FromGroup           []Resource `json:"fromGroup,omitempty"`
	FromRole            []Resource `json:"fromRole,omitempty"`
	FromSelf            []Resource `json:"fromSelf,omitempty"`
	FromResourceGallery []Resource `json:"fromResourceGallery,omitempty"`
}

type IdName struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`
}
