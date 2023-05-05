package sangforac

import (
	"encoding/json"
	"net/http"
)

type OnlineService struct {
	AC
}

// Get 获取在线用户列表,一次最多返回100个结果
func (d OnlineService) Get(filter OnlineFilter) (int, []OnlineUser, error) {
	var (
		r     = make(map[string]json.RawMessage)
		count int
		users []OnlineUser
	)
	err := d.request(http.MethodPost, "online-users").
		AddQuery("_method", "GET").
		SetBody(filter).
		Do(&r)
	if err != nil {
		return 0, nil, err
	}
	if v, ok := r["count"]; ok {
		err = json.Unmarshal(v, &count)
		if err != nil {
			return 0, nil, err
		}
	}
	if v, ok := r["users"]; ok {
		err = json.Unmarshal(v, &users)
		if err != nil {
			return 0, nil, err
		}
	}
	return count, users, nil
}

// Kick 强制注销在线用户(踢下线)
func (d OnlineService) Kick(ip string) error {
	return d.request(http.MethodPost, "online-users").
		AddQuery("_method", "DELETE").
		SetBody(map[string]string{"ip": ip}).
		Do(nil)
}

// Up 上线在线用户(单点登录)
func (d OnlineService) Up(user OnlineUser) error {
	return d.request(http.MethodPost, "online-users").
		SetBody(user).
		Do(nil)
}
