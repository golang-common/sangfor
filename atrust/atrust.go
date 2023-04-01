package atrust

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/url"
	"path"
	"time"
)

func NewAtrust(target, appid, secret string) *Atrust {
	return &Atrust{
		target: target,
		appid:  appid,
		secret: secret,
	}
}

type Atrust struct {
	target string
	appid  string
	secret string
}

func (a Atrust) Online() OnlineService {
	return OnlineService{a}
}

func (a Atrust) request(method, p string) *request {
	t := time.Now().UTC()
	return &request{
		target:    a.target,
		method:    method,
		appid:     a.appid,
		secret:    a.secret,
		uuid:      uuid.NewV4().String(),
		timestamp: fmt.Sprintf("%d", t.Unix()),
		path:      path.Join("api", p),
		query:     url.Values{},
	}
}

// IndentJson 将对象转换为更适合阅读的json格式
// 通常在调试程序时使用
func IndentJson(obj interface{}) string {
	ret, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(ret)
}

func IndentJsonBytes(b []byte) string {
	var a = make(map[string]any)
	err := json.Unmarshal(b, &a)
	if err != nil {
		return string(b)
	}
	return IndentJson(a)
}
