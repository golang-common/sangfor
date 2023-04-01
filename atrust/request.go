package atrust

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
	"strings"
	"time"
)

type request struct {
	target    string
	method    string
	appid     string
	secret    string
	uuid      string
	timestamp string
	path      string
	query     url.Values
	body      []byte
}

func (r *request) AddQuery(key, val string) *request {
	r.query.Set(key, val)
	return r
}

func (r *request) SetBody(data any) *request {
	if data != nil {
		if b, err := json.Marshal(data); err == nil {
			r.body = b
		}
	}
	return r
}

func (r *request) AddQueryData(data any) error {
	q, err := query.Values(data)
	if err != nil {
		return err
	}
	r.query = q
	return nil
}

func (r *request) Do() (RespData, error) {
	u := &url.URL{
		Scheme:   "https",
		Host:     r.target,
		Path:     r.path,
		RawQuery: r.query.Encode(),
	}
	// 构建 http 请求
	req, err := http.NewRequest(r.method, u.String(), bytes.NewReader(r.body))
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
	fmt.Println(IndentJsonBytes(res.Data))
	return res.Data, nil
}

func (r *request) addAuthHeader(h *http.Header) {
	h.Set("X-Ca-Key", r.appid)
	h.Set("X-Ca-Timestamp", r.timestamp)
	h.Set("X-Ca-Nonce", r.uuid)
	sign := r.sign()
	h.Set("X-Ca-Sign", sign)
	h.Set("Content-Type", "application/json")
}

func (r *request) sign() string {
	key := r.getSignKey()
	val := r.getSignValue()
	fmt.Printf("sign_key=\n%s\n", key)
	fmt.Printf("sign_value=\n%s\n", val)
	return r.signHmacSha256(key, val)
}

// getSignKey 获取hmac-sha256的签名key
func (r *request) getSignKey() string {
	var ul []string
	ul = append(ul, fmt.Sprintf("%s=%s", "appId", r.appid))
	ul = append(ul, fmt.Sprintf("%s=%s", "appSecret", r.secret))
	ul = append(ul, fmt.Sprintf("%s=%s", "timestamp", r.timestamp))
	ul = append(ul, fmt.Sprintf("%s=%s", "nonce", r.uuid))
	key := strings.Join(ul, "&")
	return key
}

func (r *request) getSignValue() string {
	var queryString, bodyString string
	if len(r.query) > 0 {
		queryString = r.query.Encode()
	}
	if len(r.body) > 0 {
		bodyString = string(r.body)
	}
	if queryString != "" && bodyString == "" {
		return fmt.Sprintf("/%s?%s", r.path, queryString)
	}
	if queryString == "" && bodyString != "" {
		return fmt.Sprintf("/%s?%s", r.path, bodyString)
	}
	if queryString != "" && bodyString != "" {
		return fmt.Sprintf("/%s?%s&%s", r.path, queryString, bodyString)
	}
	return "/" + r.path
}

// signHmacSha256
func (r *request) signHmacSha256(secret, msg string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(msg))
	sum := h.Sum(nil)
	sha := hex.EncodeToString(sum)
	return sha
}
