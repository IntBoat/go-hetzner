package hetzner

import (
	"os"
	"testing"
)

var client *Client
var testingServer *ServerSummary

func TestMain(m *testing.M) {
	client = NewClient(os.Getenv("username"), os.Getenv("password"))
	code := m.Run()
	os.Exit(code)
}

func TestList(t *testing.T) {
	servers, _, err := client.Server.List()
	if err != nil {
		t.Error(err)
	}

	if len(servers) == 0 {
		t.Fatal("no servers found")
	}
	t.Logf("Found %d Servers.\n", len(servers))

	testingServer = servers[0]
}

func TestGet(t *testing.T) {
	server, _, err := client.Server.Get(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
	}

	if server.ServerNumber != testingServer.ServerNumber {
		t.Error("unexpected server number")
	}

	t.Logf("ServerIP: %+v\n", server.ServerIP)
}

func TestUpdate(t *testing.T) {
	server, _, err := client.Server.Update(
		&UpdateServerRequest{
			ServerNumber: testingServer.ServerNumber,
			ServerName:   testingServer.ServerName + "_TEST",
		},
	)
	if err != nil {
		t.Error(err)
	}

	if server.ServerName != testingServer.ServerName+"_TEST" {
		t.Error("unexpected server name")
	}
	t.Logf("ServerName: %+v\n", server.ServerName)

	server, _, err = client.Server.Update(
		&UpdateServerRequest{
			ServerNumber: testingServer.ServerNumber,
			ServerName:   testingServer.ServerName,
		},
	)
	if err != nil {
		t.Error(err)
	}

	t.Logf("ServerName: %+v\n", server.ServerName)
}

func TestGetCancellation(t *testing.T) {
	server, _, err := client.Server.GetCancellation(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
	}

	t.Logf("EarliestCancellationDate: %+v\n", server.EarliestCancellationDate)
}

func TestCancelServer(t *testing.T) {
	server, _, err := client.Server.CancelServer(
		&CancelServerRequest{
			ServerNumber:       testingServer.ServerNumber,
			CancellationDate:   "2099-01-01",
			CancellationReason: "API Testing",
			ReserveLocation:    false,
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Logf("CancellationDate: %+v\n", server.CancellationDate)
}

func TestWithdrawCancellation(t *testing.T) {
	server, err := client.Server.WithdrawCancellation(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
	}

	t.Logf("StatusCode: %+v\n", server.StatusCode)
}

func TestGetCancellationAgain(t *testing.T) {
	server, _, err := client.Server.GetCancellation(testingServer.ServerNumber)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Cancelled: %+v\n", server.Cancelled)
}
