/**
 * @Author: DPY
 * @Description:
 * @File:  ac_group.go
 * @Version: 1.0.0
 * @Date: 2022/4/14 14:58
 */

package sangforac

import (
	"net/http"
)

type GroupService struct {
	AC
}

// Add 添加组
func (d GroupService) Add(group Group) error {
	err := d.request(http.MethodPost, `group`).
		SetBody(group).
		Do(nil)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除组
func (d GroupService) Delete(path string) error {
	err := d.request(http.MethodPost, `group`).
		AddQuery("_method", "DELETE").
		SetBody(Group{Path: path}).
		Do(nil)
	if err != nil {
		return err
	}
	return nil
}

// Modify 修改组
func (d GroupService) Modify(group Group) error {
	err := d.request(http.MethodPost, `group`).
		AddQuery("_method", "PUT").
		SetBody(group).
		Do(nil)
	if err != nil {
		return err
	}
	return nil
}

// SetNetPolicy 设置组的上网策略
func (d GroupService) SetNetPolicy(oper GroupPolicyOper) error {
	err := d.request(http.MethodPost, `group/netpolicy`).
		SetBody(oper).
		Do(nil)
	if err != nil {
		return err
	}
	return nil
}

// GetNetPolicy 获取组的上网策略
func (d GroupService) GetNetPolicy(path string) ([]string, error) {
	var r []string
	err := d.request(http.MethodGet, `group/netpolicy`).
		AddQuery("path", path).
		Do(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
