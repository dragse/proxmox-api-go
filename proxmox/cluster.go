package proxmox

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/client"
	error2 "github.com/dragse/proxmox-api-go/error"
	"github.com/dragse/proxmox-api-go/responses"
	"github.com/dragse/proxmox-api-go/static"
	"log"
)

type ProxmoxCluster struct {
	sessions []*client.ProxmoxSession

	Cluster *responses.ClusterStatusInformation `json:"cluster"`
	Nodes   []*responses.NodeStatusInformation  `json:"nodes"`

	currentSessionID int
}

func NewProxmoxCluster() *ProxmoxCluster {
	return &ProxmoxCluster{
		sessions:         make([]*client.ProxmoxSession, 0),
		Cluster:          nil,
		Nodes:            make([]*responses.NodeStatusInformation, 0),
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

	if tmpIndex < 0 {
		tmpIndex = 0
	}

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

func (proxmoxCluster *ProxmoxCluster) InitInformation() error {

	var respList []*json.RawMessage
	clusterStatusResponse, err := proxmoxCluster.Get(static.EndpointClusterStatus)

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
			var nodeInfo responses.NodeStatusInformation
			err = json.Unmarshal(*ele, &nodeInfo)

			if err != nil {
				return err
			}

			proxmoxCluster.Nodes = append(proxmoxCluster.Nodes, &nodeInfo)
		}
	}
	return nil
}

func (proxmoxCluster ProxmoxCluster) Get(endpoint static.Endpoint) (*responses.ProxmoxResponse, error) {
	var session *client.ProxmoxSession
	var err error

	if proxmoxCluster.currentSessionID == -1 {
		session, err = proxmoxCluster.getOnlineSession()

		if err != nil {
			return nil, err
		}
	} else {
		session = proxmoxCluster.sessions[proxmoxCluster.currentSessionID]
	}

	response, err := session.Get(endpoint)

	if err != nil {
		proxError, ok := err.(error2.SessionOfflineError)

		if !ok {
			return nil, err
		}

		log.Println(proxError.Error())
		session, err = proxmoxCluster.getOnlineSession()

		if err != nil {
			return nil, err
		}

		response, err = session.Get(endpoint)

		if err != nil {
			_, ok := err.(error2.SessionOfflineError)

			if !ok {
				return nil, err
			}

			return nil, error2.NoOnlineSessionError{SizeOfClients: len(proxmoxCluster.sessions)}
		}
	}

	return response, nil
}
