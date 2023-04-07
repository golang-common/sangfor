package test

import (
	"github.com/golang-common/sangfor/atrust/internal/model"
	"testing"
)

func TestLocalGroupSearch(t *testing.T) {
	var search = model.Search{
		UserDirectoryId: "1",
		PageSize:        4,
		PageIndex:       1,
	}
	g, err := AClient.LocalGroup().Search(search)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", g)
}
