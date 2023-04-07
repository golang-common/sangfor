package service

//
//type OnlineServ struct {
//	Atrust
//}
//
//// Search 搜索在线用户
//func (d OnlineServ) Search(search model.OnlineUserSearch) ([]model.OnlineUser, common.Pager, error) {
//	req := d.request(http.MethodGet, "v1/monitor/getUserStatus")
//	err := req.AddQueryData(search)
//	if err != nil {
//		return nil, common.Pager{}, err
//	}
//	resp, err := req.Do()
//	if err != nil {
//		return nil, common.Pager{}, err
//	}
//	var r []model.OnlineUser
//	pager, err := resp.ParseDataWithPager("data", &r)
//	if err != nil {
//		return nil, common.Pager{}, err
//	}
//	return r, pager, nil
//}
//
//// Kick 踢出在线用户，必须传入 idList 或 ulist 其中之一
//func (d OnlineServ) Kick(idList []string, ulist []model.OnlineUser) error {
//	var rm = make(map[string]any)
//	if idList == nil && ulist == nil {
//		return errors.New("give session or userdata for kick")
//	}
//	if idList != nil {
//		rm["idList"] = idList
//	}
//	if ulist != nil {
//		rm["userList"] = ulist
//	}
//	_, err := d.request(http.MethodPost, "v1/monitor/kickoutUsers").SetBody(rm).Do()
//	if err != nil {
//		return err
//	}
//	return nil
//}
