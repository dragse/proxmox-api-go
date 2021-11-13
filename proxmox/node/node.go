package node

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/client"
	"github.com/dragse/proxmox-api-go/responses/node"
	"github.com/dragse/proxmox-api-go/responses/node/vm"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"github.com/dragse/proxmox-api-go/static/timezone"
	"net/url"
)

type ProxmoxNode struct {
	client *client.ProxmoxClient

	NodeName string
}

func NewProxmoxNode(nodeName string, client *client.ProxmoxClient) *ProxmoxNode {
	return &ProxmoxNode{
		client:   client,
		NodeName: nodeName,
	}
}

func (proxmoxNode ProxmoxNode) GetNodeStatus(nodeName string) (*node.Detail, error) {
	var data *node.Detail
	resp, err := proxmoxNode.client.Get(endpoints.Nodes_Node_Status.FormatValues(nodeName))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (proxmoxNode ProxmoxNode) GetNodeTime(nodeName string) (*node.TimeInformation, error) {
	var data *node.TimeInformation
	resp, err := proxmoxNode.client.Get(endpoints.Nodes_Node_Time.FormatValues(nodeName))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (proxmoxNode ProxmoxNode) UpdateNodeTimezone(nodeName string, timezone timezone.Timezone) error {
	form := url.Values{
		"timezone": {string(timezone)},
	}

	_, err := proxmoxNode.client.PutForm(endpoints.Nodes_Node_Time.FormatValues(nodeName), form)

	if err != nil {
		return err
	}

	return nil
}

func (proxmoxNode ProxmoxNode) GetNodeVMs(nodeName string) ([]*vm.Information, error) {
	var vms []*vm.Information
	resp, err := proxmoxNode.client.Get(endpoints.Nodes_Node_Qemu.FormatValues(nodeName))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &vms)

	if err != nil {
		return nil, err
	}

	return vms, nil
}
