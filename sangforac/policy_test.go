package sangforac

import "testing"

func TestGetNetPolicy(t *testing.T) {
	npl, err := AClient.Policy().GetNetPolicy()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(IndentJson(npl))
}

func TestGetFluxPolicy(t *testing.T) {
	fpl, err := AClient.Policy().GetFluxPolicy()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(IndentJson(fpl))
}
