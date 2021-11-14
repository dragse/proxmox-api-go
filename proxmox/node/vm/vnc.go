package vm

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/responses/node/vm"
	"github.com/dragse/proxmox-api-go/static/endpoints"
)

func (proxmoxVm ProxmoxVM) CreateVNCProxy() (*vm.VNCProxy, error) {
	var url endpoints.Endpoint
	var data vm.VNCProxy

	resp, err := proxmoxVm.client.PostForm(url, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
