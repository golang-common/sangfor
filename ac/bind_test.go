package ac

import (
	"testing"
	"time"
)

var (
	ubind = UserBind{
		Name:       "dpytest",
		Enable:     true,
		Desc:       "绑定测试",
		AddrType:   "ipmac",
		Addr:       "1.1.1.1+ff-ee-ff-ee-ff-ee",
		Time:       Date{time.Now().Add(7 * 24 * time.Hour)},
		Limitlogon: true,
		Noauth: &Noauth{
			Enable:     true,
			ExpireTime: int(time.Now().Add(7 * 24 * time.Hour).Unix()),
		},
	}
	macbind = IpMacBind{
		Ip:   "2.2.2.2",
		Mac:  "dd-dd-dd-dd-dd-dd",
		Desc: "test2",
	}
)

func TestBindAddUser(t *testing.T) {
	err := AClient.Bind().AddUser(ubind)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestBindFindUser(t *testing.T) {
	rb, err := AClient.Bind().FindUser("1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(IndentJson(rb))
}

func TestBindDelUser(t *testing.T) {
	err := AClient.Bind().DelUser("1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestFindIPMac(t *testing.T) {
	r, err := AClient.Bind().FindIPMac("2.2.2.2")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(IndentJson(r))
}

func TestAddIPMac(t *testing.T) {
	err := AClient.Bind().AddIPMac(macbind)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestDelIPMac(t *testing.T) {
	err := AClient.Bind().DelIPMac("2.2.2.2")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
