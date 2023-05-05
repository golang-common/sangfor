package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Request struct {
	Target    string
	Method    string
	Appid     string
	Secret    string
	Uuid      string
	Timestamp string
	Path      string
	Query     url.Values
	Body      []byte
}

func (r *Request) AddQuery(key, val string) *Request {
	r.Query.Set(key, val)
	return r
}

func (r *Request) SetBody(data any) *Request {
	if data != nil {
		if b, err := json.Marshal(data); err == nil {
			r.Body = b
		}
	}
	return r
}

func (r *Request) AddQueryData(data any) error {
	q, err := query.Values(data)
	if err != nil {
		return err
	}
	r.Query = q
	return nil
}

func (r *Request) Do() (RespData, error) {
	u := &url.URL{
		Scheme:   "https",
		Host:     r.Target,
		Path:     r.Path,
		RawQuery: r.Query.Encode(),
	}
	// 构建 http 请求
	req, err := http.NewRequest(r.Method, u.String(), bytes.NewReader(r.Body))
	if err != nil {
		return nil, err
	}
	r.addAuthHeader(&req.Header)
	dftTsp := http.DefaultTransport.(*http.Transport)
	tsp := &http.Transport{
		Proxy:                 dftTsp.Proxy,
		DialContext:           dftTsp.DialContext,
		MaxIdleConns:          dftTsp.MaxIdleConns,
		IdleConnTimeout:       dftTsp.IdleConnTimeout,
		ExpectContinueTimeout: dftTsp.ExpectContinueTimeout,
		TLSHandshakeTimeout:   dftTsp.TLSHandshakeTimeout,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
	client := new(http.Client)
	client.Timeout = 20 * time.Second
	client.Transport = tsp
	// 发起请求
	fmt.Println(IndentJsonBytes(r.Body))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// 处理请求
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("request %s with error code %d", u.String(), resp.StatusCode))
	}
	var res Response
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return nil, err
	}
	err = res.ErrorCheck()
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (r *Request) addAuthHeader(h *http.Header) {
	h.Set("X-Ca-Key", r.Appid)
	h.Set("X-Ca-Timestamp", r.Timestamp)
	h.Set("X-Ca-Nonce", r.Uuid)
	sign := r.sign()
	h.Set("X-Ca-Sign", sign)
	h.Set("Content-Type", "application/json")
}

func (r *Request) sign() string {
	key := r.getSignKey()
	val := r.getSignValue()
	return r.signHmacSha256(key, val)
}

// getSignKey 获取hmac-sha256的签名key
func (r *Request) getSignKey() string {
	var ul []string
	ul = append(ul, fmt.Sprintf("%s=%s", "appId", r.Appid))
	ul = append(ul, fmt.Sprintf("%s=%s", "appSecret", r.Secret))
	ul = append(ul, fmt.Sprintf("%s=%s", "timestamp", r.Timestamp))
	ul = append(ul, fmt.Sprintf("%s=%s", "nonce", r.Uuid))
	key := strings.Join(ul, "&")
	return key
}

func (r *Request) getSignValue() string {
	var queryString, bodyString string
	if len(r.Query) > 0 {
		queryString = r.parseQueryString(r.Query)
	}
	if len(r.Body) > 0 {
		bodyString = string(r.Body)
	}
	if queryString != "" && bodyString == "" {
		return fmt.Sprintf("/%s?%s", r.Path, queryString)
	}
	if queryString == "" && bodyString != "" {
		return fmt.Sprintf("/%s?%s", r.Path, bodyString)
	}
	if queryString != "" && bodyString != "" {
		return fmt.Sprintf("/%s?%s&%s", r.Path, queryString, bodyString)
	}
	return "/" + r.Path
}

func (r *Request) parseQueryString(v url.Values) string {
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}

// signHmacSha256
func (r *Request) signHmacSha256(secret, msg string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(msg))
	sum := h.Sum(nil)
	sha := hex.EncodeToString(sum)
	return sha
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
