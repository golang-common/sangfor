package ac

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strings"
	"time"
)

type request struct {
	target    string
	method    string
	path      string
	secret    string
	random    string
	randomMD5 string
	query     url.Values
	data      any
	langcn    bool
	flatten   bool
}

func (r *request) SetRandom() *request {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	var (
		md5Handler = md5.New()
		rd         = fmt.Sprint(random.Uint64())
		sec        = r.secret + rd
	)
	io.WriteString(md5Handler, sec)
	r.random = rd
	r.randomMD5 = fmt.Sprintf("%x", md5Handler.Sum(nil))
	return r
}

func (r *request) SetQueryData(data any) error {
	q, err := query.Values(data)
	if err != nil {
		return err
	}
	r.query = q
	return nil
}

func (r *request) AddQuery(key, val string) *request {
	r.query.Set(key, val)
	return r
}

func (r *request) SetBody(data any) *request {
	r.data = data
	return r
}

func (r *request) SetFlatten() *request {
	r.flatten = true
	return r
}

func (r *request) SetGet() *request {
	r.method = http.MethodGet
	return r
}

func (r *request) SetPost() *request {
	r.method = http.MethodPost
	return r
}

func (r *request) Do(result any) error {
	if result != nil && reflect.TypeOf(result).Kind() != reflect.Pointer {
		return errors.New("result data must be a kind of pointer")
	}
	r.query.Add("random", r.random)
	r.query.Add("md5", r.randomMD5)
	var reqBody []byte
	if r.method == http.MethodPost {
		var temp = make(map[string]any)
		tb, err := json.Marshal(r.data)
		if err != nil {
			return err
		}
		err = json.Unmarshal(tb, &temp)
		if err != nil {
			return err
		}
		temp["random"] = r.random
		temp["md5"] = r.randomMD5
		reqBody, err = json.Marshal(temp)
		if err != nil {
			return err
		}
	}
	u := &url.URL{
		Scheme:   "http",
		Host:     r.target,
		Path:     path.Join("v1", r.path),
		RawQuery: r.query.Encode(),
	}
	// 构建 http 请求
	req, err := http.NewRequest(r.method, u.String(), bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if r.langcn {
		req.Header.Set("Accept-Language", "zh-CN")
	}
	// 构建 http 客户端
	client := new(http.Client)
	client.Timeout = 20 * time.Second
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("request %s with error code %d", u.String(), resp.StatusCode))
	}
	// 处理请求
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyBytes = fixJson(bodyBytes)
	rdata := response{}
	if result != nil {
		rdata.Data = result
	}
	err = json.Unmarshal(bodyBytes, &rdata)
	if err != nil {
		return err
	}
	err = rdata.ErrorCheck()
	if err != nil {
		return err
	}
	return nil
}

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (r response) ErrorCheck() error {
	if r.Code == 0 && r.Message != "" {
		return nil
	}
	if r.Message != "" {
		return errors.New(r.Message)
	}
	return errors.New("response check failed")
}

// fixJson 深信服AC的api接口中很多类型是字符串数组的空值[],但是返回为{},导致json解析失败
// 此方法临时做替换修复该问题
func fixJson(data []byte) []byte {
	rplcer := strings.NewReplacer(
		`"bind_cfg":{}`, `"bind_cfg":[]`,
		`"ipmac":{}`, `"ipmac":[]`,
		`"ou":{}`, `"ou":[]`,
		`"aduser":{}`, `"aduser":[]`,
		`"adgroup":{}`, `"adgroup":[]`,
		`"exc_aduser":{}`, `"exc_aduser":[]`,
		`"attribute":{}`, `"attribute":[]`,
		`"user_attr_grp":{}`, `"user_attr_grp":[]`,
		`"sourceip":{}`, `"sourceip":[]`,
		`"location":{}`, `"location":[]`,
		`"terminal":{}`, `"terminal":[]`,
		`"target_area":{}`, `"target_area":[]`,
		`"value":{}`, `"value":[]`,
	)
	return []byte(rplcer.Replace(string(data)))
}
