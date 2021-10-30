package proxmox

import "github.com/dragse/proxmox-api-go/client"

type ProxmoxCluster struct {
	sessions []*client.ProxmoxSession

	Cluster *ClusterInformation
}

type ClusterInformation struct {
	Name string
	Nodes int
	Quorate int
	Version int
}

type NodeInformation struct {
	IP string
	Level string
	Local int
	Name string
	NodeID int
	Online bool
}