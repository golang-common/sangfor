package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-common/sangfor/sangforatrust/internal/common"
	model2 "github.com/golang-common/sangfor/sangforatrust/model"
	"net/http"
)

type GroupServ struct {
	Service
}

// Search 搜索
func (d GroupServ) Search(search model2.EntityQuery) ([]model2.GroupEntity, common.Pager, error) {
	search.Type = "group"
	req := d.request(http.MethodPost, "v1/directory/queryEntity")
	req.SetBody(search)
	resp, err := req.Do()
	if err != nil {
		return nil, common.Pager{}, err
	}
	var rst []model2.GroupEntity
	err = resp.Into("group")
	if err != nil {
		return nil, common.Pager{}, err
	}
	pager, err := resp.ParseDataWithPager("data", &rst)
	if err != nil {
		return nil, common.Pager{}, err
	}
	return rst, pager, nil
}

func (d GroupServ) AddLocal(mut model2.LocalGroupAdd) error {
	req := d.request(http.MethodPost, "v2/localUserGroup/createGroup")
	req.SetBody(mut)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d GroupServ) AddExternal(mut model2.ExtGroupAdd) error {
	req := d.request(http.MethodPost, "v2/externalUserGroup/create")
	req.SetBody(mut)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d GroupServ) UpdateLocal(upd model2.LocalGroupUpdate) error {
	req := d.request(http.MethodPost, "v2/localUserGroup/updateGroup")
	req.SetBody(upd)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d GroupServ) UpdateExternal(update model2.ExtGroupUpdate) error {
	req := d.request(http.MethodPost, "v2/externalUserGroup/update")
	req.SetBody(update)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d GroupServ) DeleteLocal(idlist ...string) error {
	req := d.request(http.MethodPost, "v1/localUserGroup/deleteGroup")
	req.SetBody(map[string][]string{"idList": idlist})
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d GroupServ) GetLocalDetail(id string) (model2.LocalGroupDetail, error) {
	req := d.request(http.MethodGet, "v2/localUserGroup/queryGroup")
	req.AddQuery("id", id)
	resp, err := req.Do()
	if err != nil {
		return model2.LocalGroupDetail{}, err
	}
	var rst []model2.LocalGroupDetail
	err = resp.ParseData("data", &rst)
	if err != nil {
		return model2.LocalGroupDetail{}, err
	}
	if len(rst) == 0 {
		return model2.LocalGroupDetail{}, errors.New("not found")
	}
	return rst[0], nil
}

func (d GroupServ) GetExtDetail(arg model2.CommonArg) (model2.ExtGroupDetail, error) {
	req := d.request(http.MethodGet, "v2/externalUserGroup/queryAndGetAuthCompose")
	req.AddQueryData(arg)
	resp, err := req.Do()
	if err != nil {
		return model2.ExtGroupDetail{}, err
	}
	var rst []model2.ExtGroupDetail
	err = resp.ParseData("data", &rst)
	if err != nil {
		return model2.ExtGroupDetail{}, err
	}
	if len(rst) == 0 {
		return model2.ExtGroupDetail{}, errors.New("not found")
	}
	return rst[0], nil
}

func (d GroupServ) GetLocalDetailNested(path string) (model2.GroupNested, error) {
	req := d.request(http.MethodGet, "v1/localUserGroup/queryGroupByPath")
	req.AddQuery("path", path)
	resp, err := req.Do()
	if err != nil {
		return model2.GroupNested{}, err
	}
	var r model2.GroupNested
	err = resp.Parse(&r)
	if err != nil {
		return model2.GroupNested{}, err
	}
	return r, nil
}

func (d GroupServ) GetLocalResource(arg model2.CommonArg) (model2.ResourceRelation, error) {
	return d.getResource(arg, "local")
}

func (d GroupServ) GetExtResource(arg model2.CommonArg) (model2.ResourceRelation, error) {
	return d.getResource(arg, "external")
}

// getResource 根据用户ID或Name获取其被授权的应用和应用分类
// 返回值分别为 1-组继承的应用、2-角色继承的应用、3-用户自己直接绑定的应用、4-fromResourceGallery
func (d GroupServ) getResource(arg model2.CommonArg, src string) (model2.ResourceRelation, error) {
	var (
		res model2.ResourceRelation
	)
	req := d.request(http.MethodPost, fmt.Sprintf("v2/%sUserGroup/queryRelatedResource", src))
	req.SetBody(arg)
	resp, err := req.Do()
	if err != nil {
		return model2.ResourceRelation{}, err
	}
	var temp = map[string]map[string][]model2.Resource{}
	err = json.Unmarshal(resp, &temp)
	if err != nil {
		return model2.ResourceRelation{}, err
	}
	for k, v := range temp {
		switch k {
		case "fromGroup":
			if vd, okd := v["data"]; okd && len(vd) > 0 {
				res.FromGroup = vd
			}
		case "fromResourceGallery":
			if vd, okd := v["data"]; okd && len(vd) > 0 {
				res.FromResourceGallery = vd
			}
		case "fromRole":
			if vd, okd := v["data"]; okd && len(vd) > 0 {
				res.FromRole = vd
			}
		case "fromSelf":
			if vd, okd := v["data"]; okd && len(vd) > 0 {
				res.FromSelf = vd
			}
		}
	}
	return res, nil
}
