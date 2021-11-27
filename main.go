package main

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/client"
	"github.com/dragse/proxmox-api-go/proxmox"
	"log"
)

func main() {
	session := client.ProxmoxSession{
		Hostname:  "192.168.1.205:8006",
		Username:  "prox-api@pve!test-token",
		Token:     "1fedfb41-b8f3-40a7-8707-6f40fe617d19",
		VerifySSL: false,
	}

	proxClient := client.NewProxmoxClient()
	err := proxClient.AddSession(&session)

	if err != nil {
		log.Fatal(err)
	}

	proxCluster := proxmox.NewProxmoxCluster(proxClient)

	err = proxCluster.InitInformation()

	if err != nil {
		log.Fatal(err)
	}

	//m, err := proxClient.Get(endpoints.Pools_Pool_.FormatValues("test"))
	/*builder := builder.NewVmBuilder().
		SetID("434").
		SetName("testvm").
		SetCPUType("host").
		SetSocket(1).
		SetCoresPerSocket(3).
		SetMemory(util.NewBytesFromGigaBytes(4)).
		SetIso("local", "debian-11.0.0-amd64-netinst.iso").
		SetOSType(operation_system.L24).
		AddNetwork("vmbr0").
		AddStorage("local-lvm", "5").
		SetPool("test")
	m, err := proxCluster.GetNode("pve").CreateVM(builder)*/

	m, err := proxCluster.GetPool("test").GetDetail()

	if err != nil {
		log.Fatal(err)
	}

	test, _ := json.Marshal(m)
	log.Println(string(test))
}
