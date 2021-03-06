package vm

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/responses/node/vm"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"strconv"
)

func (proxmoxVm ProxmoxVM) CreateVNCProxy() (*vm.VNCProxy, error) {
	var data vm.VNCProxy

	resp, err := proxmoxVm.client.PostForm(endpoints.Nodes_node_Qemu_VMID_Vncproxy.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID)), nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
