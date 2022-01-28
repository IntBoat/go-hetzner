package hetzner

import "testing"

func TestResetList(t *testing.T) {
	reset, _, err := client.Reset.List()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Found %d server reset options", len(reset))
}

func TestResetGet(t *testing.T) {
	reset, _, err := client.Reset.Get(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Reset Options: %+v", reset)
}

func TestResetReset(t *testing.T) {
	//reset, _, err := client.Reset.Reset(
	//	&ResetRequest{
	//		ServerNumber: testingServer.ServerNumber,
	//		Type:         "hw",
	//	},
	//)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//t.Logf("%+v", reset)
}
