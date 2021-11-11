package main

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/client"
	"github.com/dragse/proxmox-api-go/proxmox"
	"github.com/dragse/proxmox-api-go/proxmox/builder"
	"github.com/dragse/proxmox-api-go/static/operation_system"
	"github.com/dragse/proxmox-api-go/util"
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

	//m, err := proxCluster.Get(endpoints.Nodes_Node_Qemu.FormatValues("pve"))
	builder := builder.NewVmBuilder().
		SetID("434").
		SetName("testvm").
		SetCPUType("host").
		SetSocket(1).
		SetCoresPerSocket(3).
		SetMemory(util.NewBytesFromGigaBytes(4)).
		SetIso("local", "debian-11.0.0-amd64-netinst.iso").
		SetOSType(operation_system.L24).
		AddNetwork("vmbr0").
		AddStorage("local-lvm", "32")

	m, err := proxCluster.CreateVM("pve", builder)

	if err != nil {
		log.Fatal(err)
	}

	test, _ := json.Marshal(m)
	log.Println(string(test))

}
