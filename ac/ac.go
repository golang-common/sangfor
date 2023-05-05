/**
 * @Author: DPY
 * @Description:
 * @File:  ac.go
 * @Version: 1.0.0
 * @Date: 2022/4/14 14:11
 */

package ac

import (
	"encoding/json"
	"net/url"
)

// NewAC 创建深信服AC操作对象,target为ip+端口,secret为AC上配置的密钥
// e.g: target=192.168.1.1:9999(默认端口为9999), secret=YR9nQngmvhX&9BE83K
func NewAC(target, secret string, errcn bool) *AC {
	ac := &AC{secret: secret, target: target, errLangCN: errcn}
	return ac
}

type AC struct {
	target    string
	secret    string
	errLangCN bool // 是否设置返回错误信息为中文
	debug     bool
}

func (ac AC) request(method, path string) *request {
	r := &request{
		target: ac.target,
		path:   path,
		secret: ac.secret,
		langcn: ac.errLangCN,
		method: method,
		query:  url.Values{},
	}
	return r.SetRandom()
}

func (ac AC) Status() StatusService {
	return StatusService{AC: ac}
}

func (ac AC) User() UserService {
	return UserService{AC: ac}
}

func (ac AC) Group() GroupService {
	return GroupService{AC: ac}
}

func (ac AC) Policy() PolicyService {
	return PolicyService{AC: ac}
}

func (ac AC) Online() OnlineService {
	return OnlineService{AC: ac}
}

func (ac AC) Bind() BindService {
	return BindService{AC: ac}
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
