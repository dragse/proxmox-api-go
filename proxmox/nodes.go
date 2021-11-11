package proxmox

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/responses/node"
	"github.com/dragse/proxmox-api-go/responses/node/vm"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"github.com/dragse/proxmox-api-go/static/timezone"
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

func (proxmoxCluster ProxmoxCluster) UpdateNodeTimezone(nodeName string, timezone timezone.Timezone) error {
	form := url.Values{
		"timezone": {string(timezone)},
	}

	_, err := proxmoxCluster.PutForm(endpoints.Nodes_Node_Time.FormatValues(nodeName), form)

	if err != nil {
		return err
	}

	return nil
}

func (proxmoxCluster ProxmoxCluster) GetNodeVMs(nodeName string) ([]*vm.Information, error) {
	var vms []*vm.Information
	resp, err := proxmoxCluster.Get(endpoints.Nodes_Node_Qemu.FormatValues(nodeName))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &vms)

	if err != nil {
		return nil, err
	}

	return vms, nil
}
