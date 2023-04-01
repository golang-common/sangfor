package atrust

import (
	"errors"
	"net/http"
)

type OnlineService struct {
	Atrust
}

// Search 搜索在线用户
func (d OnlineService) Search(search OnlineUserSearch) ([]OnlineUser, Pager, error) {
	req := d.request(http.MethodGet, "v1/monitor/getUserStatus")
	err := req.AddQueryData(search)
	if err != nil {
		return nil, Pager{}, err
	}
	resp, err := req.Do()
	if err != nil {
		return nil, Pager{}, err
	}
	var r []OnlineUser
	pager, err := resp.ParseDataWithPager("data", &r)
	if err != nil {
		return nil, Pager{}, err
	}
	return r, pager, nil
}

// Kick 踢出在线用户，必须传入 idList 或 ulist 其中之一
func (d OnlineService) Kick(idList []string, ulist []OnlineUser) error {
	var rm = make(map[string]any)
	if idList == nil && ulist == nil {
		return errors.New("give session or userdata for kick")
	}
	if idList != nil {
		rm["idList"] = idList
	}
	if ulist != nil {
		rm["userList"] = ulist
	}
	_, err := d.request(http.MethodPost, "v1/monitor/kickoutUsers").SetBody(rm).Do()
	if err != nil {
		return err
	}
	return nil
}
