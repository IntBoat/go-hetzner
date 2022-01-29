package hetzner

import (
	"strings"
	"testing"
)

// No money to order Windows, cPanel and Plesk to test :(

func TestGetBoot(t *testing.T) {
	server, _, err := client.Boot.GetBoot(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	var boot []string
	if server.Rescue != nil {
		boot = append(boot, "Rescue")
	}
	if server.Rescue != nil {
		boot = append(boot, "Linux")
	}
	if server.Rescue != nil {
		boot = append(boot, "Rescue")
	}

	if len(boot) == 0 {
		t.Error("No boot supported.")
		return
	}

	t.Log("Supported boot: " + strings.Join(boot, ", "))
}

func TestGetRescue(t *testing.T) {
	server, _, err := client.Boot.GetRescue(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}
	if server == nil {
		t.Error("Rescue boot not supported.")
		return
	}

	//t.Logf("Boot: %+v", server)
}

func TestGetRescueLast(t *testing.T) {
	server, _, err := client.Boot.GetRescueLast(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	if server == nil {
		t.Error("Rescue boot not supported.")
		return
	}

	//t.Logf("Boot: %+v", server)
}

func TestGetLinux(t *testing.T) {
	server, _, err := client.Boot.GetLinux(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	if server == nil {
		t.Error("Linux boot not supported.")
		return
	}

	//t.Logf("Boot: %+v", server)
}

func TestGetLinuxLast(t *testing.T) {
	server, _, err := client.Boot.GetLinuxLast(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	if server == nil {
		t.Error("Linux boot not supported.")
		return
	}

	//t.Logf("Boot: %+v", server)
}

func TestGetVnc(t *testing.T) {
	server, _, err := client.Boot.GetVnc(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	if server == nil {
		t.Error("VNC boot not supported.")
		return
	}

	//t.Logf("Boot: %+v", server)
}

func TestGetWindows(t *testing.T) {
	server, _, err := client.Boot.GetWindows(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	if server == nil {
		t.Error("Windows boot not supported.")
		return
	}

	//t.Logf("Boot: %+v", server)
}

func TestGetPlesk(t *testing.T) {
	server, _, err := client.Boot.GetPlesk(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	if server == nil {
		t.Error("Plesk boot not supported.")
		return
	}

	//t.Logf("Boot: %+v", server)
}

func TestGetCPanel(t *testing.T) {
	server, _, err := client.Boot.GetCPanel(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	if server == nil {
		t.Error("CPanel boot not supported.")
		return
	}

	//t.Logf("Boot: %+v", server)
}

func TestActivateRescue(t *testing.T) {
	server, _, err := client.Boot.ActivateRescue(
		&ActivateRescueRequest{
			ServerNumber:  testingServer.ServerNumber,
			OS:            "linux",
			Arch:          64,
			AuthorizedKey: nil,
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	if server.Active != true {
		t.Error("Can not activate boot mode")
		return
	}

	t.Logf("Boot: %+v", server.Active)
}

func TestDeactivateRescue(t *testing.T) {
	server, _, err := client.Boot.DeactivateRescue(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	if server.Active != false {
		t.Error("Can not deactivate boot mode")
		return
	}

	t.Logf("Boot: %+v", server.Active)
}

func TestActivateLinux(t *testing.T) {
	server, _, err := client.Boot.ActivateLinux(
		&ActivateLinuxRequest{
			ServerNumber:  testingServer.ServerNumber,
			Dist:          "Debian 10.11 minimal",
			Arch:          64,
			Lang:          "en",
			AuthorizedKey: nil,
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	if server.Active != true {
		t.Error("Can not activate boot mode")
		return
	}

	t.Logf("Boot: %+v", server.Active)
}

func TestDeactivateLinux(t *testing.T) {
	server, _, err := client.Boot.DeactivateLinux(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	if server.Active != false {
		t.Error("Can not deactivate boot mode")
		return
	}

	t.Logf("Boot: %+v", server.Active)
}

func TestActivateVnc(t *testing.T) {
	server, _, err := client.Boot.ActivateVnc(
		&ActivateVncRequest{
			ServerNumber: testingServer.ServerNumber,
			Dist:         "CentOS-7.9",
			Arch:         64,
			Lang:         "en_US",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	if server.Active != true {
		t.Error("Can not activate boot mode")
		return
	}

	t.Logf("Boot: %+v", server.Active)
}

func TestDeactivateVnc(t *testing.T) {
	server, _, err := client.Boot.DeactivateVnc(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}

	if server.Active != false {
		t.Error("Can not deactivate boot mode")
		return
	}

	t.Logf("Boot: %+v", server.Active)
}

func TestActivateWindows(t *testing.T) {
	server, _, err := client.Boot.ActivateWindows(
		&ActivateWindowsRequest{
			ServerNumber: testingServer.ServerNumber,
			Lang:         "en",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	if server.Active != true {
		t.Error("Can not activate boot mode")
		return
	}

	t.Logf("Boot: %+v", server.Active)
}

func TestDeactivateWindows(t *testing.T) {
	server, _, err := client.Boot.DeactivateWindows(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}
	if server.Active != false {
		t.Error("Can not deactivate boot mode")
		return
	}

	t.Logf("Boot: %+v", server.Active)
}

func TestActivatePlesk(t *testing.T) {
	hostname := "API Testing"
	server, _, err := client.Boot.ActivatePlesk(
		&ActivatePleskRequest{
			ServerNumber: testingServer.ServerNumber,
			Dist:         "Debian 7.8 minimal",
			Arch:         64,
			Lang:         "en",
			Hostname:     &hostname,
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	if server.Active != true {
		t.Error("Can not activate boot mode")
		return
	}

	t.Logf("Boot: %+v", server.Active)
}

func TestDeactivatePlesk(t *testing.T) {
	server, _, err := client.Boot.DeactivatePlesk(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}
	if server.Active != false {
		t.Error("Can not deactivate boot mode")
		return
	}

	t.Logf("Boot: %+v", server)
}

func TestActivateCPanel(t *testing.T) {
	hostname := "API Testing"
	server, _, err := client.Boot.ActivateCPanel(
		&ActivateCPanelRequest{
			ServerNumber: testingServer.ServerNumber,
			Dist:         "Debian 10.11 minimal",
			Arch:         64,
			Lang:         "en",
			Hostname:     &hostname,
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	if server.Active != true {
		t.Error("Can not activate boot mode")
		return
	}

	t.Logf("Boot: %+v", server.Active)
}

func TestDeactivateCPanel(t *testing.T) {
	server, _, err := client.Boot.DeactivateCPanel(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
		return
	}
	if server.Active != false {
		t.Error("Can not deactivate boot mode")
		return
	}

	t.Logf("Boot: %+v", server)
}
