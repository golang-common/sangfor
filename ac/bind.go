package ac

import (
	"net/http"
)

type BindService struct {
	AC
}

// FindUser 查询用户和IP/MAC的绑定关系(支持按用户名,IP,MAC进行搜索)
func (d BindService) FindUser(search string) ([]UserBind, error) {
	var r []UserBind
	err := d.request(http.MethodGet, "bindinfo/user-bindinfo").
		AddQuery("search", search).
		Do(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// AddUser 增加用户和IP/MAC的绑定关系
func (d BindService) AddUser(bind UserBind) error {
	return d.request(http.MethodPost, "bindinfo/user-bindinfo").
		SetBody(bind).
		Do(nil)
}

// DelUser 删除用户和IP/MAC的绑定关系,addr可以是IP或MAC地址
func (d BindService) DelUser(addr string) error {
	return d.request(http.MethodPost, "bindinfo/user-bindinfo").
		AddQuery("_method", "DELETE").
		SetBody(map[string]string{"addr": addr}).
		Do(nil)
}

// FindIPMac 查询IP/Mac绑定关系(支持按ip/mac进行搜索)
func (d BindService) FindIPMac(search string) (IpMacBind, error) {
	var r IpMacBind
	err := d.request(http.MethodGet, "ipmac-bindinfo").
		AddQuery("search", search).
		Do(&r)
	if err != nil {
		return IpMacBind{}, err
	}
	return r, nil
}

// AddIPMac 增加IP/Mac的绑定关系
func (d BindService) AddIPMac(bind IpMacBind) error {
	return d.request(http.MethodPost, "bindinfo/ipmac-bindinfo").
		SetBody(bind).
		Do(nil)
}

// DelIPMac 删除IP/MAC的绑定关系
func (d BindService) DelIPMac(ip string) error {
	return d.request(http.MethodPost, "bindinfo/ipmac-bindinfo").
		AddQuery("_method", "DELETE").
		SetBody(map[string]string{"ip": ip}).
		Do(nil)
}
