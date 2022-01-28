package hetzner

import "testing"

var rdnsPtr = ""

func TestRDNSList(t *testing.T) {
	rdns, _, err := client.RDNS.List()
	if err != nil {
		t.Error(err)
		return
	}

	if len(rdns) == 0 {
		t.Fatal("No RDNS found")
		return
	}

	t.Logf("Found %d RDNS records", len(rdns))
}

func TestRDNSGet(t *testing.T) {
	rdns, _, err := client.RDNS.Get(testingServer.ServerIP)
	if err != nil {
		t.Error(err)
		return
	}

	rdnsPtr = rdns.Ptr
	t.Logf("RDNS: %+v", rdns)
}

func TestRDNSDelete(t *testing.T) {
	resp, err := client.RDNS.Delete(testingServer.ServerIP)
	if err != nil && resp.StatusCode != 200 {
		t.Error(err)
		return
	}

	if resp.StatusCode != 200 {
		t.Error("Failed to remove rDNS entry")
	}
}

func TestRDNSCreate(t *testing.T) {
	rdns, resp, err := client.RDNS.Create(
		&RDNSUpdateRequest{
			IP:  testingServer.ServerIP,
			Ptr: rdnsPtr,
		},
	)
	if err != nil && resp.StatusCode != 409 {
		t.Error(err)
		return
	}

	t.Logf("RDNS: %+v", rdns)
}

func TestRDNSUpdate(t *testing.T) {
	rdns, _, err := client.RDNS.Update(
		&RDNSUpdateRequest{
			IP:  testingServer.ServerIP,
			Ptr: rdnsPtr,
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("RDNS: %+v", rdns)
}
