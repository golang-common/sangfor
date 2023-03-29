package ac

import "testing"

func TestAddGroup(t *testing.T) {
	err := AClient.Group().Add(Group{
		Path: "/dpytest",
		Desc: "戴澎源的测试组",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestDelGroup(t *testing.T) {
	err := AClient.Group().Delete("/dpytest")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestModifyGroup(t *testing.T) {
	err := AClient.Group().Modify(Group{
		Path: "/dpytest",
		Desc: "戴澎源的测试组改",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestGroupGetNetPolicy(t *testing.T) {
	pl, err := AClient.Group().GetNetPolicy("/dpytest")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pl)
}

func TestGroupSetNetPolicy(t *testing.T) {
	err := AClient.Group().SetNetPolicy(GroupPolicyOper{
		Group:  "/dpytest",
		Opr:    "del",
		Policy: []string{"腾讯文档-白名单"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
