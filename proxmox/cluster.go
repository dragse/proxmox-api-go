package proxmox

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/client"
	error2 "github.com/dragse/proxmox-api-go/error"
	"github.com/dragse/proxmox-api-go/responses"
	"github.com/dragse/proxmox-api-go/responses/node"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"log"
	"net/url"
)

type ProxmoxCluster struct {
	sessions []*client.ProxmoxSession

	Cluster *responses.ClusterStatusInformation `json:"cluster"`
	Nodes   []*node.StatusInformation           `json:"nodes"`

	currentSessionID int
}

func NewProxmoxCluster() *ProxmoxCluster {
	return &ProxmoxCluster{
		sessions:         make([]*client.ProxmoxSession, 0),
		Cluster:          nil,
		Nodes:            make([]*node.StatusInformation, 0),
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
	clusterStatusResponse, err := proxmoxCluster.Get(endpoints.ClusterStatus)

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

func (proxmoxCluster ProxmoxCluster) PostForm(endpoint endpoints.Endpoint, form url.Values) (*responses.ProxmoxResponse, error) {
	return proxmoxCluster.execute(func(session *client.ProxmoxSession) (*responses.ProxmoxResponse, error) {
		return session.PostForm(endpoint, form)
	})
}

func (proxmoxCluster ProxmoxCluster) PutForm(endpoint endpoints.Endpoint, form url.Values) (*responses.ProxmoxResponse, error) {
	return proxmoxCluster.execute(func(session *client.ProxmoxSession) (*responses.ProxmoxResponse, error) {
		return session.PutForm(endpoint, form)
	})
}

func (proxmoxCluster ProxmoxCluster) Get(endpoint endpoints.Endpoint) (*responses.ProxmoxResponse, error) {
	return proxmoxCluster.execute(func(session *client.ProxmoxSession) (*responses.ProxmoxResponse, error) {
		return session.Get(endpoint)
	})
}

func (proxmoxCluster ProxmoxCluster) execute(data func(session *client.ProxmoxSession) (*responses.ProxmoxResponse, error)) (*responses.ProxmoxResponse, error) {
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

	response, err := data(session)

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

		response, err = data(session)

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
