package test

import (
	"fmt"
	"github.com/golang-common/sangfor/atrust/model"
	"testing"
	"time"
)

func TestLocalGroupSearch(t *testing.T) {
	var search = model.EntityQuery{
		UserDirectoryId: "1",
		PageSize:        4,
		PageIndex:       1,
	}
	g, p, err := AClient.Group().Search(search)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", g)
	t.Log(p)
}

func TestExtUserResourceGet(t *testing.T) {
	rl, err := AClient.User().GetExtResource(model.CommonArg{
		Name:              "DaiPengYuan",
		UserDirectoryName: "兴盛优选_企业微信扫码",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(IndentJson(rl))
}

func TestNewLocalUser(t *testing.T) {
	tm := time.Now().UnixMilli()
	fmt.Println(tm)
	add := model.LocalUserAdd{}
	add.Name = "lyonsdpy"
	add.DisplayName = "戴澎源"
	add.GroupId = "19b7cc00-d773-11ed-8d5d-b7730edbdea3"
	add.Phone = "13070167676"
	add.Email = "lyonsdpy@163.com"
	add.ExpiredTime = fmt.Sprintf("%d", tm)
	add.Password = "z6f*i7w{0`2ipcz"
	add.PwdModel = "clear"
	err := AClient.User().AddLocal(add)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestGetGroup(t *testing.T) {
	eq := model.EntityQuery{}
	eq.UserDirectoryId = "1"
	eq.PageSize = 20
	eq.PageIndex = 1
	eq.Path = "/"
	ett, pg, err := AClient.Group().Search(eq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pg)
	t.Log(ett)
}

func TestGetDetail(t *testing.T) {
	n, err := AClient.Group().GetLocalDetail("19b7cc00-d773-11ed-8d5d-b7730edbdea3")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(n)
}

func TestGetNestGroup(t *testing.T) {
	n, err := AClient.Group().GetLocalDetailNested("/tmp")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(n)
}
