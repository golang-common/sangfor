package model

type Role struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ExtRoleDetail struct {
	RoleDetail
	Type            int    `json:"type,omitempty"`
	UserDirectoryId string `json:"userDirectoryId,omitempty"`
}

type RoleDetail struct {
	Role
	ResourceList []Resource `json:"resourceList,omitempty"`
}

type LocalRoleAdd struct {
	RoleAdd
	ResourceIdList      []string `json:"resourceIdList,omitempty"`
	ResourceGroupIdList []string `json:"resourceGroupIdList,omitempty"`
}

type ExtRoleAdd struct {
	RoleAdd
	UserDirectoryId   string   `json:"userDirectoryId,omitempty"`
	UserDirectoryName string   `json:"userDirectoryName,omitempty"`
	UserNameList      []string `json:"userNameList,omitempty"`
}

type RoleAdd struct {
	Role
	UserIdList    []string `json:"userIdList,omitempty"`
	GroupIdList   []string `json:"groupIdList,omitempty"`
	GroupNameList []string `json:"groupNameList,omitempty"`
}

type ExtRoleUpdate struct {
	UserDirectoryId   string `json:"userDirectoryId,omitempty"`
	UserDirectoryName string `json:"userDirectoryName,omitempty"`
	RoleUpdate
}

type RoleUpdate struct {
	Role
	UpdateUser
	UpdateResource
	UpdateResourceGroup
	UpdateGroup
}

type RoleEntity struct {
	Role
	UserDirectoryId string `json:"userDirectoryId,omitempty"`
	Type            int    `json:"type,omitempty"`
}
