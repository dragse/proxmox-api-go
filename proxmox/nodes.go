package proxmox

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/responses/node"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"net/url"
)

func (proxmoxCluster ProxmoxCluster) GetNodes() ([]*node.Information, error) {
	var nodes []*node.Information
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

func (proxmoxCluster ProxmoxCluster) GetNodeStatus(nodeName string) (*node.Detail, error) {
	var data *node.Detail
	resp, err := proxmoxCluster.Get(endpoints.Nodes_Node_Status.FormatValues(nodeName))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (proxmoxCluster ProxmoxCluster) GetNodeTime(nodeName string) (*node.TimeInformation, error) {
	var data *node.TimeInformation
	resp, err := proxmoxCluster.Get(endpoints.Nodes_Node_Time.FormatValues(nodeName))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (proxmoxCluster ProxmoxCluster) UpdateNodeTimezone(nodeName string, timezone string) error {
	form := url.Values{
		"timezone": {timezone},
	}

	_, err := proxmoxCluster.PutForm(endpoints.Nodes_Node_Time.FormatValues(nodeName), form)

	if err != nil {
		return err
	}

	return nil
}
