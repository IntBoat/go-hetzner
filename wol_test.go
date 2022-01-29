package hetzner

import "testing"

func TestWolGet(t *testing.T) {
	wol, _, err := client.Wol.Get(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Wol: %+v", wol)
}

func TestWolCreate(t *testing.T) {
	wol, _, err := client.Wol.Create(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Wol: %+v", wol)
}
