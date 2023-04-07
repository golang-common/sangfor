package atrust

import "github.com/golang-common/sangfor/atrust/internal/service"

func NewAtrust(target, appid, secret string) *service.Atrust {
	return &service.Atrust{
		Target: target,
		Appid:  appid,
		Secret: secret,
	}
}
