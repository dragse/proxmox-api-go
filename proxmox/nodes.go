package proxmox

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/responses"
	"github.com/dragse/proxmox-api-go/static/endpoints"
)

func (proxmoxCluster ProxmoxCluster) GetNodes() ([]*responses.NodeInformation, error) {
	var nodes []*responses.NodeInformation
	resp, err := proxmoxCluster.Get(endpoints.Nodes)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &nodes)

	if err != nil {
		return nil, err
	}

	return nodes, nil
}
