package ac

import (
	"testing"
)

var ou = OnlineUser{
	Name:       "dpytest2",
	ShowName:   "戴澎源测试2",
	FatherPath: `/兴盛优选/五区经营管委会/技术中心/供应链&运营支撑研发部`,
	Ip:         "10.10.10.20",
	Mac:        "ff-ff-ff-ff-ee-ff",
}

func TestOnlineGET(t *testing.T) {
	count, users, err := AClient.Online().Get(OnlineFilter{FilterType: "user", FilterValue: []string{"dpytest2"}})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
	t.Log(len(users))
}

func TestOnlineUp(t *testing.T) {
	t.Log(IndentJson(ou))
	err := AClient.Online().Up(ou)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("success")
}

func TestOnlineKick(t *testing.T) {
	err := AClient.Online().Kick("10.10.10.20")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
