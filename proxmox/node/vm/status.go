package vm

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/responses/node/vm"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"github.com/dragse/proxmox-api-go/static/status"
	"strconv"
)

func (proxmoxVm ProxmoxVM) GetVMStatus() (*vm.Detail, error) {
	var data vm.Detail
	resp, err := proxmoxVm.client.Get(endpoints.Nodes_Node_Qemu_VMID_StatusCurrent.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID)))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (proxmoxVm ProxmoxVM) UpdateVMStatus(updateStatus status.Status) (string, error) {
	var url endpoints.Endpoint
	var data string

	switch updateStatus {
	case status.Reboot:
		url = endpoints.Nodes_Node_Qemu_VMID_StatusReboot.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID))
	case status.Reset:
		url = endpoints.Nodes_Node_Qemu_VMID_StatusReset.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID))
	case status.Resume:
		url = endpoints.Nodes_Node_Qemu_VMID_StatusResume.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID))
	case status.Shutdown:
		url = endpoints.Nodes_Node_Qemu_VMID_StatusShutdown.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID))
	case status.Start:
		url = endpoints.Nodes_Node_Qemu_VMID_StatusStart.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID))
	case status.Stop:
		url = endpoints.Nodes_Node_Qemu_VMID_StatusStop.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID))
	case status.Suspend:
		url = endpoints.Nodes_Node_Qemu_VMID_StatusSuspend.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID))
	}

	resp, err := proxmoxVm.client.PostForm(url, nil)

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return "", err
	}

	return data, nil
}
