package vm

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/proxmox/builder"
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

func (proxmoxVm ProxmoxVM) GetFirewallOptions() (*vm.FirewallOption, error) {
	var data vm.FirewallOption

	resp, err := proxmoxVm.client.Get(endpoints.Nodes_Node_Qemu_VMID_FirewallOptions.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID)))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (proxmoxVm ProxmoxVM) UpdateFirewallOptions(firewallBuilder *builder.FirewallBuilder) error {
	_, err := proxmoxVm.client.PostForm(endpoints.Nodes_Node_Qemu_VMID_FirewallOptions.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID)), firewallBuilder.BuildToValues())

	if err != nil {
		return err
	}

	return nil
}

func (proxmoxVm ProxmoxVM) ListIPSets() ([]*vm.IPSet, error) {
	var data []*vm.IPSet

	resp, err := proxmoxVm.client.Get(endpoints.Nodes_Node_Qemu_VMID_FirewallIpset.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID)))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
