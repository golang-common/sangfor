/**
 * @Author: DPY
 * @Description:
 * @File:  ac_user.go
 * @Version: 1.0.0
 * @Date: 2022/4/14 14:38
 */

package ac

import (
	"net/http"
)

type UserService struct {
	AC
}

// Add 添加新用户
func (d UserService) Add(data UserAdd) error {
	var r string
	err := d.request(http.MethodPost, `user`).
		SetBody(data).
		Do(&r)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除用户
func (d UserService) Delete(username string) error {
	dt := UserAdd{}
	dt.Name = username
	err := d.request(http.MethodPost, `user`).
		AddQuery("_method", "DELETE").
		SetBody(dt).
		Do(nil)
	if err != nil {
		return err
	}
	return nil
}

// Modify 修改用户
func (d UserService) Modify(data UserModify) error {
	err := d.request(http.MethodPost, `user`).
		AddQuery("_method", "PUT").
		SetBody(data).
		Do(nil)
	if err != nil {
		return err
	}
	return nil
}

// Search 搜索用户
func (d UserService) Search(data UserSearch) ([]UserDetail, error) {
	var r []UserDetail
	err := d.request(http.MethodPost, `user`).
		AddQuery("_method", "GET").
		SetBody(data).
		Do(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Get 获取用户详细信息
func (d UserService) Get(username string) (UserDetail, error) {
	var r UserDetail
	err := d.request(http.MethodGet, `user`).
		AddQuery("name", username).
		Do(&r)
	if err != nil {
		return UserDetail{}, err
	}
	return r, nil
}

// GetNetPolicy 获取用户的上网策略列表
func (d UserService) GetNetPolicy(username string, dn ...string) ([]string, error) {
	var r []string
	req := d.request(http.MethodGet, `user/netpolicy`)
	req.AddQuery("user", username)
	if len(dn) > 0 {
		req.AddQuery("dn", dn[0])
	}
	err := req.Do(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// SetNetPolicy 设置用户的上网策略
func (d UserService) SetNetPolicy(oper UserPolicyOper) error {
	var r string
	err := d.request(http.MethodPost, `user/netpolicy`).
		SetBody(oper).
		Do(&r)
	if err != nil {
		return err
	}
	return nil
}

// GetFluxPolicy 获取用户的流控策略
func (d UserService) GetFluxPolicy(username string, dn ...string) ([]FluxPolicy, error) {
	var r []FluxPolicy
	req := d.request(http.MethodGet, `user/fluxpolicy`)
	req.AddQuery("user", username)
	if len(dn) > 0 {
		req.AddQuery("dn", dn[0])
	}
	err := req.Do(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// SetFluxPolicy 设置用户的流控策略
func (d UserService) SetFluxPolicy(oper UserPolicyOper) error {
	err := d.request(http.MethodPost, `user/fluxpolicy`).
		SetBody(oper).
		Do(nil)
	if err != nil {
		return err
	}
	return nil
}

// VerifyPwd 验证用户密码
// TODO: 该接口目前实际效果是获取用户
func (d UserService) VerifyPwd(username, password string) error {
	err := d.request(http.MethodGet, `user`).
		AddQuery("_method", "verify").
		AddQuery("name", username).
		AddQuery("password", password).
		Do(nil)
	if err != nil {
		return err
	}
	return nil
}
