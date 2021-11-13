package vm

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/client"
	"github.com/dragse/proxmox-api-go/responses/node/vm"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"strconv"
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

func (proxmoxCluster ProxmoxVM) GetVMStatus() (*vm.Detail, error) {
	var data vm.Detail
	resp, err := proxmoxCluster.client.Get(endpoints.Nodes_Node_Qemu_VMID_StatusCurrent.FormatValues(proxmoxCluster.NodeName, strconv.Itoa(proxmoxCluster.VmID)))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
