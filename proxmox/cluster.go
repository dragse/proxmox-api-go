package proxmox

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/client"
	"github.com/dragse/proxmox-api-go/responses"
	"github.com/dragse/proxmox-api-go/responses/node"
	"github.com/dragse/proxmox-api-go/static/endpoints"
)

type ProxmoxCluster struct {
	client *client.ProxmoxClient

	Cluster *responses.ClusterStatusInformation `json:"cluster"`
	Nodes   []*node.StatusInformation           `json:"nodes"`

	currentSessionID int
}

func NewProxmoxCluster(proxmoxClient *client.ProxmoxClient) *ProxmoxCluster {
	return &ProxmoxCluster{
		client:  proxmoxClient,
		Cluster: nil,
		Nodes:   make([]*node.StatusInformation, 0),
	}
}

func (proxmoxCluster *ProxmoxCluster) InitInformation() error {

	var respList []*json.RawMessage
	clusterStatusResponse, err := proxmoxCluster.client.Get(endpoints.ClusterStatus)

	if err != nil {
		return err
	}

	err = json.Unmarshal(*clusterStatusResponse.Data, &respList)

	if err != nil {
		return err
	}

	type respElement struct {
		Type string `json:"type"`
	}

	for _, ele := range respList {
		var resp respElement
		err = json.Unmarshal(*ele, &resp)

		if err != nil {
			return err
		}

		switch resp.Type {
		case "cluster":
			var clusterInfo responses.ClusterStatusInformation
			err = json.Unmarshal(*ele, &clusterInfo)

			if err != nil {
				return err
			}

			proxmoxCluster.Cluster = &clusterInfo
		case "node":
			var nodeInfo node.StatusInformation
			err = json.Unmarshal(*ele, &nodeInfo)

			if err != nil {
				return err
			}

			proxmoxCluster.Nodes = append(proxmoxCluster.Nodes, &nodeInfo)
		}
	}
	return nil
}
