package service

import (
	"fmt"
	"github.com/golang-common/sangfor/atrust/internal/common"
	uuid "github.com/satori/go.uuid"
	"net/url"
	"path"
	"time"
)

type Atrust struct {
	Target string
	Appid  string
	Secret string
}

//func (a Atrust) Online() OnlineServ {
//	return OnlineServ{a}
//}

func (a Atrust) User() UserServ {
	return UserServ{a}
}

func (a Atrust) Group() GroupServ {
	return GroupServ{a}
}

func (a Atrust) Role() RoleServ {
	return RoleServ{a}
}

func (a Atrust) request(method, p string) *common.Request {
	t := time.Now().UTC()
	return &common.Request{
		Target:    a.Target,
		Method:    method,
		Appid:     a.Appid,
		Secret:    a.Secret,
		Uuid:      uuid.NewV4().String(),
		Timestamp: fmt.Sprintf("%d", t.Unix()),
		Path:      path.Join("api", p),
		Query:     url.Values{},
	}
}
