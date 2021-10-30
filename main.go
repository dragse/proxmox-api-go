package main

import (
	"github.com/dragse/proxmox-api-go/client"
	"log"
)

func main() {
	hist := client.ProxmoxSession{
		Hostname:  "192.168.1.205:8006",
		Username:  "prox-api@pve!test-token",
		Token:  "1fedfb41-b8f3-40a7-8707-6f40fe617d19",
		VerifySSL: false,
	}

	err := hist.SetupClient()

	if err != nil {
		log.Fatal(err)
	}

	err = hist.TestConnection()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Login Success")
}
