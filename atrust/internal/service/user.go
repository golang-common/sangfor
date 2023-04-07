package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-common/sangfor/atrust/internal/common"
	"github.com/golang-common/sangfor/atrust/internal/model"
	"net/http"
)

type UserServ struct {
	Atrust
}

// Search 搜索用户
func (d UserServ) Search(search model.EntityQuery) ([]model.UserEntity, common.Pager, error) {
	search.Type = "user"
	req := d.request(http.MethodPost, "v2/directory/queryEntity")
	req.SetBody(search)
	resp, err := req.Do()
	if err != nil {
		return nil, common.Pager{}, err
	}
	var rst []model.UserEntity
	err = resp.Into("user")
	if err != nil {
		return nil, common.Pager{}, err
	}
	pager, err := resp.ParseDataWithPager("data", &rst)
	if err != nil {
		return nil, common.Pager{}, err
	}
	return rst, pager, nil
}

// GetLocalDetail 查询本地用户详情（支持根据name查询）
func (d UserServ) GetLocalDetail(query model.CommonArg) (model.LocalUserDetail, error) {
	req := d.request(http.MethodGet, "v2/localUser/queryUser")
	req.AddQueryData(query)
	resp, err := req.Do()
	if err != nil {
		return model.LocalUserDetail{}, err
	}
	var r []model.LocalUserDetail
	err = resp.ParseData("data", &r)
	if err != nil {
		return model.LocalUserDetail{}, err
	}
	if len(r) == 0 {
		return model.LocalUserDetail{}, errors.New("not found")
	}
	return r[0], nil
}

func (d UserServ) GetExtDetail(query model.CommonArg) (model.ExtUserDetail, error) {
	req := d.request(http.MethodGet, "v2/externalUser/queryUser")
	req.AddQueryData(query)
	resp, err := req.Do()
	if err != nil {
		return model.ExtUserDetail{}, err
	}
	var r []model.ExtUserDetail
	err = resp.ParseData("data", &r)
	if err != nil {
		return model.ExtUserDetail{}, err
	}
	if len(r) == 0 {
		return model.ExtUserDetail{}, errors.New("not found")
	}
	return r[0], nil
}

// AddLocal 新增本地用户
func (d UserServ) AddLocal(user model.LocalUserAdd) error {
	req := d.request(http.MethodPost, "v2/localUser/createUser")
	req.SetBody(user)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d UserServ) AddExternal(user model.ExtUserAdd) error {
	req := d.request(http.MethodPost, "v2/externalUser/create")
	req.SetBody(user)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

// UpdateLocal 修改本地用户
func (d UserServ) UpdateLocal(mut model.LocalUserUpdate) error {
	req := d.request(http.MethodPost, "v2/localUser/updateUser")
	req.SetBody(mut)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d UserServ) UpdateExternal(mut model.ExtUserUpdate) error {
	req := d.request(http.MethodPost, "v2/externalUser/update")
	req.SetBody(mut)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

// BatchUpdateLocal 批量修改本地用户
func (d UserServ) BatchUpdateLocal(idlist []string, attr model.UserUpdate) error {
	var temp = make(map[string]any)
	temp["idList"] = idlist
	temp["value"] = attr
	req := d.request(http.MethodPost, "v2/localUser/batchUpdateUser")
	req.SetBody(temp)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d UserServ) BatchDeleteExt(arg model.CommonArg, attr model.ExtUserBatchUpdate) error {
	arg.Value = attr
	req := d.request(http.MethodPost, "v2/externalUser/batchUpdate")
	req.SetBody(arg)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d UserServ) DeleteLocal(idlist ...string) error {
	var temp = map[string][]string{"idList": idlist}
	req := d.request(http.MethodPost, "v1/localUser/batchDeleteUser")
	req.SetBody(temp)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d UserServ) DeleteExternal(arg model.CommonArg) error {
	req := d.request(http.MethodPost, "v1/externalUser/delete")
	req.SetBody(arg)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d UserServ) GetLocalResource(arg model.CommonArg) (model.ResourceRelation, error) {
	return d.getResource(arg, "local")
}

func (d UserServ) GetExtResource(arg model.CommonArg) (model.ResourceRelation, error) {
	return d.getResource(arg, "external")
}

// getResource 根据用户ID或Name获取其被授权的应用和应用分类
// 返回值分别为 1-组继承的应用、2-角色继承的应用、3-用户自己直接绑定的应用、4-fromResourceGallery
func (d UserServ) getResource(arg model.CommonArg, src string) (model.ResourceRelation, error) {
	req := d.request(http.MethodPost, fmt.Sprintf("v2/%sUser/queryRelatedResource", src))
	req.SetBody(arg)
	resp, err := req.Do()
	if err != nil {
		return model.ResourceRelation{}, err
	}
	return d.parseResource(resp)
}

func (d UserServ) parseResource(resp common.RespData) (model.ResourceRelation, error) {
	var res model.ResourceRelation
	var temp = map[string]map[string][]model.Resource{}
	err := json.Unmarshal(resp, &temp)
	if err != nil {
		return model.ResourceRelation{}, err
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
