package vm

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/proxmox/builder"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"strconv"
)

func (proxmoxVm ProxmoxVM) Clone(builder *builder.VmCopyBuilder) (string, error) {
	var data string

	resp, err := proxmoxVm.client.PostForm(endpoints.Nodes_node_Qemu_VMID_Clone.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID)), builder.BuildToValues())

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return "", err
	}

	return data, nil
}
