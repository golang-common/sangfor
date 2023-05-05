package model

import (
	"encoding/json"
)

// EntityQuery 搜索参数
// Type - 搜索类型,user,group
type EntityQuery struct {
	Type              string   `json:"type,omitempty"`
	UserDirectoryId   string   `json:"userDirectoryId,omitempty"`
	UserDirectoryName string   `json:"userDirectoryName,omitempty"`
	ShowAll           int      `json:"showAll,omitempty"`
	PageSize          int      `json:"pageSize,omitempty"`
	PageIndex         int      `json:"pageIndex"`
	Path              string   `json:"path,omitempty"`
	SearchData        []string `json:"searchData,omitempty"`
}

func (u EntityQuery) MarshalJSON() ([]byte, error) {
	var s = make(map[string]any)
	if u.UserDirectoryId != "" {
		s["userDirectoryId"] = u.UserDirectoryId
	}
	if u.UserDirectoryName != "" {
		s["userDirectoryName"] = u.UserDirectoryName
	}
	s["showAll"] = u.ShowAll
	if u.PageSize != 0 && u.PageIndex != 0 && u.Type != "" {
		s["include"] = map[string]map[string]int{
			u.Type: {
				"pageSize":  u.PageSize,
				"pageIndex": u.PageIndex,
			},
		}
	}
	if u.Path != "" {
		s["path"] = u.Path
	}
	if len(u.SearchData) > 0 {
		s["searchData"] = u.SearchData
	}
	r, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type EntityAttr struct {
	Creator    string `json:"creator,omitempty"`
	DataType   string `json:"dataType,omitempty"`
	ServerName string `json:"serverName,omitempty"`
	EntityType string `json:"entityType,omitempty"`
	Domain     string `json:"domain,omitempty"`
}
