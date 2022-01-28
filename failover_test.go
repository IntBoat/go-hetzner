package hetzner

import "testing"

// No money to order a failover ip to test :(

func TestFailoverList(t *testing.T) {
	failover, resp, err := client.Failover.List()
	if err != nil && resp.StatusCode != 404 {
		t.Error(err)
		return
	}

	if resp.StatusCode == 404 || len(failover) == 0 {
		t.Error("No failover IP addresses found")
		return
	}

	t.Logf("IP: %+v\n", failover[0].IP)
}

func TestFailoverGet(t *testing.T) {
	failover, resp, err := client.Failover.Get("0.0.0.0")
	if err != nil && resp.StatusCode != 404 {
		t.Error(err)
		return
	}

	if resp.StatusCode == 404 {
		t.Error("No failover IP addresses found")
		return
	}

	t.Logf("IP: %+v\n", failover.IP)
}

func TestFailoverSwitch(t *testing.T) {
	failover, resp, err := client.Failover.Switch(
		FailoverSwitchRequest{
			FailoverIP:     "0.0.0.0",
			ActiveServerIP: "0.0.0.0",
		},
	)
	if err != nil && resp.StatusCode != 404 {
		t.Error(err)
		return
	}

	if resp.StatusCode == 404 {
		t.Error("No failover IP addresses found")
		return
	}

	t.Logf("IP: %+v\n", failover.IP)
}

func TestFailoverDelete(t *testing.T) {
	failover, resp, err := client.Failover.Delete("0.0.0.0")
	if err != nil && resp.StatusCode != 404 {
		t.Error(err)
		return
	}

	if resp.StatusCode == 404 {
		t.Error("No failover IP addresses found")
		return
	}

	t.Logf("IP: %+v\n", failover.IP)
}
