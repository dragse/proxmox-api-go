package main

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/client"
	"github.com/dragse/proxmox-api-go/proxmox"
	"github.com/dragse/proxmox-api-go/static"
	"log"
)

func main() {
	session := client.ProxmoxSession{
		Hostname:  "192.168.1.205:8006",
		Username:  "prox-api@pve!test-token",
		Token:     "1fedfb41-b8f3-40a7-8707-6f40fe617d19",
		VerifySSL: false,
	}

	proxCluster := proxmox.NewProxmoxCluster()

	err := proxCluster.AddSession(&session)

	if err != nil {
		log.Fatal(err)
	}

	err = proxCluster.InitInformation()

	if err != nil {
		log.Fatal(err)
	}

	proxCluster.Get(static.EndpointClusterStatus)

	test, _ := json.Marshal(proxCluster)
	log.Println(string(test))

}
