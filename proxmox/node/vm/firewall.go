package vm

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/responses/node/vm"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"strconv"
)

func (proxmoxVm ProxmoxVM) GetFirewallLog() ([]*vm.FirewallLog, error) {
	var data []*vm.FirewallLog

	resp, err := proxmoxVm.client.Get(endpoints.Nodes_Node_Qemu_VMID_FirewallLog.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID)))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
