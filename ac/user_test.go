package ac

import (
	"testing"
	"time"
)

var uex = UserExtend[[]UserBindCfg, []string]{
	Enable: true,
	SelfPass: &UserSelfPass{
		Enable:     true,
		Password:   "daipengyuan",
		ModifyOnce: true,
	},
	BindCfg: []UserBindCfg{
		{
			Ip:       "1.1.1.1",
			Mac:      "ff-ff-ff-ff-ff-ff",
			OutTime:  Date{time.Now()},
			Bindgoal: "noauth",
			Desc:     "testbind1",
		},
		{
			Ip:       "2.2.2.2",
			Mac:      "ee-ee-ee-ee-ee-ee",
			OutTime:  Date{time.Now()},
			Bindgoal: "loginlimit",
			Desc:     "testbind2",
		},
	},
	CustomCfg: []map[string]string{
		{"customkey1": "customval1"},
		{"customkey2": "customval2"},
	},
	Logout: true,
	CommonUser: &UserCommon{
		AllowChange: true,
		Enable:      true,
	},
	LimitIpmac: []string{
		"192.168.1.1-192.168.1.3",
		"192.168.1.2",
	},
}

func TestAddUser(t *testing.T) {
	user := UserAdd{
		UserBasic: UserBasic[Time]{
			Name:       "dpytest",
			Desc:       "戴澎源描述",
			ShowName:   "戴澎源显示名",
			FatherPath: "/兴盛优选",
			ExpireTime: Time{time.Now().Add(365 * 24 * time.Hour)},
		},
		UserExtend: uex,
	}
	t.Log(user.ExpireTime.String())
	t.Log(IndentJson(user))
	err := AClient.User().Add(user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestDeleteUser(t *testing.T) {
	err := AClient.User().Delete("dpytest")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestGetUser(t *testing.T) {
	ud, err := AClient.User().Get("dpytest")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(IndentJson(ud))
}

func TestModifyUser(t *testing.T) {
	mod := UserModify{Name: "dpytest"}
	mod.Desc = "戴澎源描述改"
	mod.LimitIpmac = &UserLimit{
		Enable: true,
		Ipmac:  []string{"192.168.1.1"},
	}
	mod.ExpireTime = Time{Time: time.Now().Add(180 * 24 * time.Hour)}
	err := AClient.User().Modify(mod)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestSearchUser(t *testing.T) {
	ud, err := AClient.User().Search(UserSearch{
		SearchType:  "user",
		SearchValue: "dpytest",
		Extend:      nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(IndentJson(ud))
}

func TestGetUserNetPolicy(t *testing.T) {
	pl, err := AClient.User().GetNetPolicy("dpytest")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pl)
}

func TestSetUserNetPolicy(t *testing.T) {
	err := AClient.User().SetNetPolicy(UserPolicyOper{
		User:   "dpytest",
		Opr:    "add",
		Policy: []string{"语雀-白名单"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestGetUserFluxPolicy(t *testing.T) {
	pl, err := AClient.User().GetFluxPolicy("dpytest")
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range pl {
		t.Log(p.Name)
	}
}

func TestSetUserFluxPolicy(t *testing.T) {
	err := AClient.User().SetFluxPolicy(UserPolicyOper{
		User:   "dpytest",
		Opr:    "del",
		Policy: []string{"流量优先保障"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestVerifyPwd(t *testing.T) {
	err := AClient.User().VerifyPwd("dpytest", "daipengyua2n")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
