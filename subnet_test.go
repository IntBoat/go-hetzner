package hetzner

import "testing"

// No money to order a subnet to test :(
func TestSubnetList(t *testing.T) {
	subnet, resp, err := client.Subnet.List()
	if err != nil && resp.StatusCode != 404 {
		t.Error(err)
		return
	}

	if resp.StatusCode == 404 {
		t.Error("No IP addresses found")
		return
	}

	t.Logf("Found %d IPs", len(subnet))
}

func TestSubnetGet(t *testing.T) {
	subnet, resp, err := client.Subnet.Get(testingServer.ServerIP)
	if err != nil && resp.StatusCode != 404 {
		t.Error(err)
		return
	}

	if resp.StatusCode == 404 {
		t.Error("No IP addresses found")
		return
	}

	t.Logf("Subnet: %+v\n", subnet.IP)
}

func TestSubnetUpdateWarring(t *testing.T) {
	subnet, resp, err := client.Subnet.UpdateWarring(
		&SubnetUpdateRequest{
			NetIP:           testingServer.ServerIP,
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

	if subnet.TrafficWarnings != true {
		t.Error("Update subnet traffic warning failed")
		return
	}

	subnet, resp, err = client.Subnet.UpdateWarring(
		&SubnetUpdateRequest{
			NetIP:           testingServer.ServerIP,
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

	if subnet.TrafficWarnings != false {
		t.Error("Update subnet traffic warning failed")
		return
	}
}

func TestSubnetGetMac(t *testing.T) {
	subnet, _, err := client.Subnet.GetMac(testingServer.ServerIP)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Subnet Mac: %+v\n", subnet.Mac)
}

func TestSubnetCreateMac(t *testing.T) {
	subnet, _, err := client.Subnet.CreateMac(testingServer.ServerIP)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Subnet Mac: %+v\n", subnet.Mac)
}

func TestSubnetDeleteMac(t *testing.T) {
	subnet, _, err := client.Subnet.DeleteMac(testingServer.ServerIP)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Subnet Mac: %+v\n", subnet.Mac)
}
