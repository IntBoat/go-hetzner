package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IntBoat/go-hetzner"
)

func main() {
	if len(os.Args) != 3 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage:\n\t%s username password\n", os.Args[0])
		os.Exit(1)
	}

	username, password := os.Args[1], os.Args[2]

	client := hetzner.NewClient(username, password)
	servers, _, err := client.Server.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, srv := range servers {
		fmt.Printf("%s\t%s\n", srv.ServerName, srv.ServerIP)
	}
}
