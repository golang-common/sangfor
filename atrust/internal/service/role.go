package service

import (
	"errors"
	"github.com/golang-common/sangfor/atrust/internal/common"
	model2 "github.com/golang-common/sangfor/atrust/model"
	"net/http"
)

type RoleServ struct {
	Service
}

func (d RoleServ) SearchExternal(query model2.EntityQuery) ([]model2.RoleEntity, common.Pager, error) {
	query.Type = "role"
	req := d.request(http.MethodPost, "v2/directory/queryEntity")
	req.SetBody(query)
	resp, err := req.Do()
	if err != nil {
		return nil, common.Pager{}, err
	}
	var rst []model2.RoleEntity
	err = resp.Into("role")
	if err != nil {
		return nil, common.Pager{}, err
	}
	pager, err := resp.ParseDataWithPager("data", &rst)
	if err != nil {
		return nil, common.Pager{}, err
	}
	return rst, pager, nil
}

func (d RoleServ) AddLocal(add model2.LocalRoleAdd) error {
	req := d.request(http.MethodPost, "v2/localUserRole/create")
	req.SetBody(add)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d RoleServ) AddExternal(add model2.ExtRoleAdd) error {
	req := d.request(http.MethodPost, "v2/externalUserRole/create")
	req.SetBody(add)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d RoleServ) UpdateLocal(update model2.RoleUpdate) error {
	req := d.request(http.MethodPost, "v2/localUserRole/update")
	req.SetBody(update)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d RoleServ) UpdateExternal(update model2.ExtRoleUpdate) error {
	req := d.request(http.MethodPost, "v2/externalUserRole/update")
	req.SetBody(update)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d RoleServ) DeleteLocal(idlist ...string) error {
	req := d.request(http.MethodPost, "v2/localUserRole/delete")
	req.SetBody(map[string][]string{"idList": idlist})
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d RoleServ) DeleteExternal(arg model2.CommonArg) error {
	req := d.request(http.MethodPost, "v2/externalUserRole/delete")
	req.SetBody(arg)
	_, err := req.Do()
	if err != nil {
		return err
	}
	return nil
}

func (d RoleServ) GetLocalDetail(arg model2.CommonArg) (model2.RoleDetail, error) {
	req := d.request(http.MethodGet, "v2/localUserRole/queryById")
	req.AddQueryData(arg)
	resp, err := req.Do()
	if err != nil {
		return model2.RoleDetail{}, err
	}
	var r []model2.RoleDetail
	err = resp.ParseData("data", &r)
	if err != nil {
		return model2.RoleDetail{}, err
	}
	if len(r) == 0 {
		return model2.RoleDetail{}, errors.New("not found")
	}
	return r[0], nil
}

func (d RoleServ) GetExtDetail(arg model2.CommonArg) (model2.ExtRoleDetail, error) {
	req := d.request(http.MethodGet, "v2/externalUserRole/queryById")
	req.AddQueryData(arg)
	resp, err := req.Do()
	if err != nil {
		return model2.ExtRoleDetail{}, err
	}
	var r []model2.ExtRoleDetail
	err = resp.ParseData("data", &r)
	if err != nil {
		return model2.ExtRoleDetail{}, err
	}
	if len(r) == 0 {
		return model2.ExtRoleDetail{}, errors.New("not found")
	}
	return r[0], nil
}
