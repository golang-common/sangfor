package atrust

import "testing"

func TestOnlineService_Search(t *testing.T) {
	var search = OnlineUserSearch{
		PageSize:    50,
		PageIndex:   1,
		Filter:      "name",
		SearchValue: "XuYong",
	}
	ol, pg, err := AClient.Online().Search(search)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", pg)
	t.Log()
	t.Logf(IndentJson(ol))
}
