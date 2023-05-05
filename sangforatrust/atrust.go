package sangforatrust

import "github.com/golang-common/sangfor/sangforatrust/internal/service"

func NewAtrust(target, appid, secret string) *Atrust {
	return &Atrust{
		Target: target,
		Appid:  appid,
		Secret: secret,
	}
}

type Atrust struct {
	Target string
	Appid  string
	Secret string
}

func (a Atrust) service() service.Service {
	return service.Service{
		Target: a.Target,
		Appid:  a.Appid,
		Secret: a.Secret,
	}
}

//func (a Atrust) Online() OnlineServ {
//	return OnlineServ{a}
//}

func (a Atrust) User() service.UserServ {
	return service.UserServ{Service: a.service()}
}

func (a Atrust) Group() service.GroupServ {
	return service.GroupServ{Service: a.service()}
}

func (a Atrust) Role() service.RoleServ {
	return service.RoleServ{Service: a.service()}
}
