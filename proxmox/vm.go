package proxmox

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/proxmox/builder"
	"github.com/dragse/proxmox-api-go/static/endpoints"
)

func (proxmoxCluster ProxmoxCluster) CreateVM(nodeName string, builder *builder.VmBuilder) (string, error) {
	var data string
	resp, err := proxmoxCluster.PostForm(endpoints.Nodes_Node_Qemu.FormatValues(nodeName), builder.BuildToValues())

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return "", err
	}

	return data, nil
}
