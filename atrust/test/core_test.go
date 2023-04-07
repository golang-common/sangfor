package test

import (
	"encoding/json"
	"github.com/golang-common/sangfor/atrust"
	"testing"
	"time"
)

var AClient = atrust.NewAtrust("10.9.255.2:4433",
	"1461600",
	"6023d77f29144216bcac2496208872c8")

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

func TestTime(t *testing.T) {
	cl, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatal(err)
	}
	tm, err := time.ParseInLocation(`2006-01-02 15:04:05`, `2021-08-21 14:25:00`, cl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tm.Unix())

	tm, err = time.Parse(`2006-01-02 15:04:05`, `2021-08-21 14:25:00`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tm.Unix() - 8*3600)
}

type Ta struct {
	A string `json:"a,omitempty"`
	B string `json:"b,omitempty"`
}

type Tb struct {
	C string `json:"c,omitempty"`
	D string `json:"d,omitempty"`
}

func TestJSON(t *testing.T) {
	var a = `{"a":"a","b":"b","c":"c","d":"d"}`
	var ta Ta
	var tb Tb
	var err error
	err = json.Unmarshal([]byte(a), &ta)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", ta)
	err = json.Unmarshal([]byte(a), &tb)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", tb)
}
