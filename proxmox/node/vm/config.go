package vm

import (
	"github.com/dragse/proxmox-api-go/proxmox/builder"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"strconv"
)

func (proxmoxVm ProxmoxVM) UpdateConfigASync(builder *builder.VmBuilder) error {
	_, err := proxmoxVm.client.PostForm(endpoints.Nodes_node_Qemu_VMID_Config.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID)), builder.BuildToValues())

	if err != nil {
		return err
	}

	return nil
}
