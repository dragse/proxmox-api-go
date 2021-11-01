package main

import (
	"github.com/dragse/proxmox-api-go/client"
	"github.com/dragse/proxmox-api-go/proxmox"
	"github.com/dragse/proxmox-api-go/static/timezone"
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

	//m, err := proxCluster.Get(endpoints.Nodes_Node_Time.FormatValues("pve"))
	err = proxCluster.UpdateNodeTimezone("pve", timezone.Europe_Berlin)

	if err != nil {
		log.Fatal(err)
	}

	//test, _ := json.Marshal(m)
	//log.Println(string(test))

}
