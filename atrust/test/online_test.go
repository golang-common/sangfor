package test

import (
	"github.com/golang-common/sangfor/atrust/internal/model"
	"testing"
)

func TestOnlineService_Search(t *testing.T) {
	var search = model.OnlineUserSearch{
		PageSize:    10,
		PageIndex:   1,
		Filter:      "name",
		SearchValue: "wangbinbin",
	}
	ol, pg, err := AClient.Online().Search(search)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", pg)
	t.Log()
	t.Logf(IndentJson(ol))
}
