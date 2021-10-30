package proxmox

import "github.com/dragse/proxmox-api-go/client"

type ProxmoxCluster struct {
	sessions []*client.ProxmoxSession

	Cluster *ClusterInformation
	Nodes   []*NodeInformation
}

type ClusterInformation struct {
	Name    string `json:"name"`
	Nodes   int    `json:"nodes"`
	Quorate int    `json:"quorate"`
	Version int    `json:"version"`
}

type NodeInformation struct {
	IP     string `json:"ip"`
	Level  string `json:"level"`
	Local  int    `json:"local"`
	Name   string `json:"name"`
	NodeID int    `json:"node_id"`
	Online bool   `json:"online"`
}
