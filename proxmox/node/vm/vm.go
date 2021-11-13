package vm

import (
	"github.com/dragse/proxmox-api-go/client"
)

type ProxmoxVM struct {
	client *client.ProxmoxClient

	NodeName string
	VmID     int
}

func NewProxmoxVM(client *client.ProxmoxClient, nodeName string, vmID int) *ProxmoxVM {
	return &ProxmoxVM{
		client:   client,
		NodeName: nodeName,
		VmID:     vmID,
	}
}
