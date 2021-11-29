package vm

import (
	"github.com/dragse/proxmox-api-go/static/disk"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"github.com/dragse/proxmox-api-go/util"
	"net/url"
	"strconv"
)

func (proxmoxVm ProxmoxVM) Resize(disk disk.DiskType, size *util.Byte) error {
	values := url.Values{}

	values.Add("disk", string(disk))
	values.Add("size", strconv.FormatInt(size.ToGigaByte(), 10)+"G")

	_, err := proxmoxVm.client.PutForm(endpoints.Nodes_node_Qemu_VMID_Resize.FormatValues(proxmoxVm.NodeName, strconv.Itoa(proxmoxVm.VmID)), values)

	if err != nil {
		return err
	}

	return nil
}
