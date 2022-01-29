package hetzner

import "testing"

// No money to order a new ip to test :(
func TestIPList(t *testing.T) {
	IPs, resp, err := client.IP.List()
	if err != nil && resp.StatusCode != 404 {
		t.Error(err)
		return
	}

	if resp.StatusCode == 404 {
		t.Error("No IP addresses found")
		return
	}

	t.Logf("Found %d IPs", len(IPs))
}

func TestIPGet(t *testing.T) {
	IPs, resp, err := client.IP.Get(testingServer.ServerIP)
	if err != nil && resp.StatusCode != 404 {
		t.Error(err)
		return
	}

	if resp.StatusCode == 404 {
		t.Error("IP address not found")
		return
	}

	t.Logf("IP: %+v\n", IPs.ServerIP)
}

func TestIPUpdateWarring(t *testing.T) {
	IPs, resp, err := client.IP.UpdateWarring(
		&IPUpdateRequest{
			IP:              testingServer.ServerIP,
			TrafficWarnings: "true",
			TrafficHourly:   "100",
			TrafficDaily:    "100",
			TrafficMonthly:  "100",
		},
	)
	if err != nil && resp.StatusCode != 404 {
		t.Error(err)
		return
	}

	if resp.StatusCode == 404 {
		t.Error("IP address not found")
		return
	}

	if IPs.TrafficWarnings != true {
		t.Error("Update traffic warning failed")
		return
	}

	IPs, resp, err = client.IP.UpdateWarring(
		&IPUpdateRequest{
			IP:              testingServer.ServerIP,
			TrafficWarnings: "false",
		},
	)
	if err != nil && resp.StatusCode != 404 {
		t.Error(err)
		return
	}

	if resp.StatusCode == 404 {
		t.Error("IP address not found")
		return
	}

	if IPs.TrafficWarnings != false {
		t.Error("Update traffic warning failed")
		return
	}
}

func TestIPGetMac(t *testing.T) {
	IPs, _, err := client.IP.GetMac(testingServer.ServerIP)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("IP Mac: %+v\n", IPs.Mac)
}

func TestIPCreateMac(t *testing.T) {
	IPs, _, err := client.IP.CreateMac(testingServer.ServerIP)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("IP Mac: %+v\n", IPs.Mac)
}

func TestIPDeleteMac(t *testing.T) {
	IPs, _, err := client.IP.DeleteMac(testingServer.ServerIP)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("IP Mac: %+v\n", IPs.Mac)
}
