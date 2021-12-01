package main

import (
	"github.com/dragse/proxmox-api-go/client"
	"github.com/dragse/proxmox-api-go/proxmox"
	"github.com/dragse/proxmox-api-go/proxmox/builder"
	"github.com/dragse/proxmox-api-go/static/disk"
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

	//m, err := proxCluster.GetPool("test").GetDetail()
	_, err = proxCluster.GetNode("pve").GetVM(105).Clone(builder.NewVmCopyBuilder().SetFullCopy(true).SetPool("test").SetName("copy").SetTargetNode("pve").SetNewID(999))

	if err != nil {

		log.Fatal(err)
	}

	copyBuilder := builder.NewVmBuilder().
		SetName("cloudCoyp").
		SetCoresPerSocket(4).
		SetMemory(util.NewBytesFromGigaBytes(16)).
		SetCloudInitDrive("local-lvm").
		SetCloudInit(builder.NewCloudInitBuilder().
			SetNameservers("1.1.1.1").
			SetUser("hans").
			AddIPConfig("10.0.0.10/24", "10.0.0.1", "", ""))

	vals := copyBuilder.BuildToValues()

	for k, v := range vals {
		log.Println(k, " => ", v)
	}

	err = proxCluster.GetNode("pve").GetVM(999).UpdateConfigASync(copyBuilder)

	if err != nil {
		log.Fatal(err)
	}

	err = proxCluster.GetNode("pve").GetVM(999).Resize(disk.Scsi0, util.NewBytesFromGigaBytes(20))

	if err != nil {
		log.Fatal(err)
	}

	//test, _ := json.Marshal(m)
	//log.Println(string(test))
}
