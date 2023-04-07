package model

type Group struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type LocalGroupAdd struct {
	GroupAdd
	ParentGroupId string `json:"parentGroupId"`
}

type ExtGroupAdd struct {
	GroupAdd
	UserDirectoryId   string `json:"userDirectoryId,omitempty"`
	UserDirectoryName string `json:"userDirectoryName,omitempty"`
}

type GroupAdd struct {
	Group
	RoleIdList      []string `json:"roleIdList,omitempty"`
	RoleNameList    []string `json:"roleNameList,omitempty"`
	UserPolicyId    string   `json:"userPolicyId,omitempty"`
	UserPolicyName  string   `json:"userPolicyName,omitempty"`
	AuthComposeId   string   `json:"authComposeId,omitempty"`
	AuthComposeName string   `json:"authComposeName,omitempty"`
}

type LocalGroupUpdate struct {
	GroupUpdate
	ParentGroupId string `json:"parentGroupId,omitempty"`
}

type ExtGroupUpdate struct {
	GroupUpdate
	UserDirectoryId   string `json:"userDirectoryId,omitempty"`
	UserDirectoryName string `json:"userDirectoryName,omitempty"`
	Path              string `json:"path,omitempty"`
}

type GroupUpdate struct {
	Group
	UpdateRole
	UpdateResource
	UpdateResourceGroup
	AuthComposeId   string `json:"authComposeId,omitempty"`
	AuthComposeName string `json:"authComposeName,omitempty"`
	UserPolicyId    string `json:"userPolicyId,omitempty"`
	UserPolicyName  string `json:"userPolicyName,omitempty"`
}

type GroupEntity struct {
	Group
	EntityAttr
	Path string `json:"path,omitempty"`
}

type LocalGroupDetail struct {
	Group
	ParentGroupId   string `json:"parentGroupId,omitempty"`
	ParentGroupName string `json:"parentGroupName,omitempty"`
	Depth           string `json:"depth,omitempty"`
	AuthCompose     IdName `json:"authCompose,omitempty"`
	UserPolicy      IdName `json:"userPolicy,omitempty"`
}

type ExtGroupDetail struct {
	Group
	UserDirectoryId string `json:"userDirectoryId,omitempty"`
	Status          int    `json:"status,omitempty"`
	ExpiredTime     string `json:"expiredTime,omitempty"`
	UserPolicyId    string `json:"userPolicyId,omitempty"`
	AuthComposeId   string `json:"authComposeId,omitempty"`
	AuthComposeName string `json:"authComposeName,omitempty"`
}

type GroupDetail struct {
	Group
	Path     string `json:"path,omitempty"`
	RoleList IdName `json:"roleList,omitempty"`
}

type GroupNested struct {
	Group
	ParentGroupId string        `json:"parentGroupId,omitempty"`
	HasChild      int           `json:"hasChild,omitempty"`
	Path          string        `json:"path,omitempty"`
	Children      []GroupNested `json:"children,omitempty"`
}
