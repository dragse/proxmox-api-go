package node

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/proxmox/builder"
	"github.com/dragse/proxmox-api-go/proxmox/node/vm"
	"github.com/dragse/proxmox-api-go/static/endpoints"
)

func (proxmoxNode ProxmoxNode) GetVM(vmID int) *vm.ProxmoxVM {
	return vm.NewProxmoxVM(proxmoxNode.client, proxmoxNode.NodeName, vmID)
}

func (proxmoxNode ProxmoxNode) CreateVM(builder *builder.VmBuilder) (string, error) {
	var data string
	resp, err := proxmoxNode.client.PostForm(endpoints.Nodes_Node_Qemu.FormatValues(proxmoxNode.NodeName), builder.BuildToValues())

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return "", err
	}

	return data, nil
}
