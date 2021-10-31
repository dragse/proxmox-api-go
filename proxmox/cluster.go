package proxmox

import (
	"github.com/dragse/proxmox-api-go/client"
	error2 "github.com/dragse/proxmox-api-go/error"
)

type ProxmoxCluster struct {
	sessions []*client.ProxmoxSession

	Cluster *ClusterInformation
	Nodes   []*NodeInformation

	currentSessionID int
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

func NewProxmoxCluster() *ProxmoxCluster {
	return &ProxmoxCluster{
		sessions:         make([]*client.ProxmoxSession, 0),
		Cluster:          nil,
		Nodes:            make([]*NodeInformation, 0),
		currentSessionID: -1,
	}
}

func (proxmoxCluster ProxmoxCluster) AddSessionFromValues(hostname string, username string, token string, verifySSL bool) error {
	return proxmoxCluster.AddSession(&client.ProxmoxSession{
		Hostname:  hostname,
		Username:  username,
		Token:     token,
		VerifySSL: verifySSL,
	})
}

func (proxmoxCluster *ProxmoxCluster) AddSession(session *client.ProxmoxSession) error {
	err := session.SetupClient()

	if err != nil {
		return err
	}

	proxmoxCluster.sessions = append(proxmoxCluster.sessions, session)
	return nil
}

func (proxmoxCluster *ProxmoxCluster) getOnlineSession() (*client.ProxmoxSession, error) {
	tmpIndex := proxmoxCluster.currentSessionID

	for true {

		err := proxmoxCluster.sessions[tmpIndex].TestConnection()

		if err == nil {
			break
		}

		tmpIndex++

		if tmpIndex >= len(proxmoxCluster.sessions) {
			tmpIndex = 0
		}

		if proxmoxCluster.currentSessionID == tmpIndex {
			// all sessions checked and none is online
			return nil, error2.NoOnlineSessionError{
				SizeOfClients: len(proxmoxCluster.sessions),
			}
		}
	}

	proxmoxCluster.currentSessionID = tmpIndex
	return proxmoxCluster.sessions[tmpIndex], nil
}
