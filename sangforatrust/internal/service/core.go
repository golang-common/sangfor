package service

import (
	"fmt"
	"github.com/golang-common/sangfor/sangforatrust/internal/common"
	uuid "github.com/satori/go.uuid"
	"net/url"
	"path"
	"time"
)

type Service struct {
	Target string
	Appid  string
	Secret string
}

func (a Service) request(method, p string) *common.Request {
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
