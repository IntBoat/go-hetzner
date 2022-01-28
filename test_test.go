package hetzner

import (
	"fmt"
	"os"
	"testing"
)

var client *Client
var testingServer *ServerSummary

// Fetch 1 server for testing
func TestMain(m *testing.M) {
	client = NewClient(os.Getenv("username"), os.Getenv("password"))
	servers, _, err := client.Server.List()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
		return
	}

	if len(servers) == 0 {
		fmt.Print("no servers found")
		os.Exit(1)
		return
	}

	testingServer = servers[0]
	code := m.Run()
	os.Exit(code)
}
