package proxmox

import (
	"encoding/json"
	node2 "github.com/dragse/proxmox-api-go/proxmox/node"
	"github.com/dragse/proxmox-api-go/responses/node"
	"github.com/dragse/proxmox-api-go/static/endpoints"
)

func (proxmoxCluster ProxmoxCluster) GetNodes() ([]*node.Information, error) {
	var nodes []*node.Information
	resp, err := proxmoxCluster.client.Get(endpoints.Nodes)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &nodes)

	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func (proxmoxCluster ProxmoxCluster) GetNode(nodeName string) *node2.ProxmoxNode {
	return node2.NewProxmoxNode(nodeName, proxmoxCluster.client)
}
