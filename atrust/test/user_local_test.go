package test

import (
	"github.com/golang-common/sangfor/atrust/internal/model"
	"testing"
)

func TestUserService_GetLocalDetail(t *testing.T) {
	resp, err := AClient.LocalUser().GetByName("yangjingxing")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(IndentJson(resp))
}

func TestUserService_Search(t *testing.T) {
	resp, pager, err := AClient.LocalUser().Search(model.Search{
		UserDirectoryId: "1",
		PageSize:        1,
		PageIndex:       1,
		//Path:            "/内部人员",
		SearchData: []string{"杨"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", pager)
	t.Log(IndentJson(resp))
}

func TestUserService_Update(t *testing.T) {
	var mdata = model.LocalUserMut{}
	mdata.Name = "DaiPengyuan"
	mdata.Phone = ""
	err := AClient.LocalUser().Update(mdata)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestUserService_Add(t *testing.T) {
	var add = model.LocalUserMut{
		UserBase: model.UserBase{
			Name:        "DaiPengyuan",
			Description: "戴澎源描述",
			DisplayName: "戴澎源描述",
			Phone:       "13070167676",
			Email:       "lyonsdpy@163.com",
			Status:      1,
		},
		UserMutAttr: model.UserMutAttr{
			RoleEditWay:  "set",
			RoleNameList: []string{"DBA手机账号密码登录"},
		},
		Password: "Dpy,./123456",
		PwdModel: "clear",
		GroupId:  "root",
	}
	err := AClient.LocalUser().Add(add)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestUserService_BatchUpdate(t *testing.T) {
	err := AClient.LocalUser().BatchUpdate([]string{"1152cd40-7fda-11ec-9f89-952e7d605ff6"}, model.UserMutAttr{
		RoleEditWay:  "set",
		RoleNameList: []string{"DBA手机账号密码登录"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestUserService_BatchDelete(t *testing.T) {
	err := AClient.LocalUser().BatachDelete([]string{"7c5dacb0-d441-11ed-8d5d-b7730edbdea3"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestUserService_GetResourceByName(t *testing.T) {
	f, err := AClient.LocalUser().GetResourceByName("yangjingxing")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(IndentJson(f))
}

func TestUserService_GetResourceByID(t *testing.T) {
	f, err := AClient.LocalUser().GetResourceByID("1152cd40-7fda-11ec-9f89-952e7d605ff6")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(IndentJson(f))
}
