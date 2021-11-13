package client

import (
	error2 "github.com/dragse/proxmox-api-go/error"
	"github.com/dragse/proxmox-api-go/responses"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"log"
	"net/url"
)

type ProxmoxClient struct {
	sessions         []*ProxmoxSession
	currentSessionID int
}

func NewProxmoxClient() *ProxmoxClient {
	return &ProxmoxClient{
		sessions:         make([]*ProxmoxSession, 0),
		currentSessionID: -1,
	}
}

func (proxmoxClient ProxmoxClient) AddSessionFromValues(hostname string, username string, token string, verifySSL bool) error {
	return proxmoxClient.AddSession(&ProxmoxSession{
		Hostname:  hostname,
		Username:  username,
		Token:     token,
		VerifySSL: verifySSL,
	})
}

func (proxmoxClient *ProxmoxClient) AddSession(session *ProxmoxSession) error {
	err := session.SetupClient()

	if err != nil {
		return err
	}

	proxmoxClient.sessions = append(proxmoxClient.sessions, session)
	return nil
}

func (proxmoxClient *ProxmoxClient) getOnlineSession() (*ProxmoxSession, error) {
	tmpIndex := proxmoxClient.currentSessionID

	if tmpIndex < 0 {
		tmpIndex = 0
	}

	for true {

		err := proxmoxClient.sessions[tmpIndex].TestConnection()

		if err == nil {
			break
		}

		tmpIndex++

		if tmpIndex >= len(proxmoxClient.sessions) {
			tmpIndex = 0
		}

		if proxmoxClient.currentSessionID == tmpIndex {
			// all sessions checked and none is online
			return nil, error2.NoOnlineSessionError{
				SizeOfClients: len(proxmoxClient.sessions),
			}
		}
	}

	proxmoxClient.currentSessionID = tmpIndex
	return proxmoxClient.sessions[tmpIndex], nil
}

func (proxmoxClient ProxmoxClient) PostForm(endpoint endpoints.Endpoint, form url.Values) (*responses.ProxmoxResponse, error) {
	return proxmoxClient.execute(func(session *ProxmoxSession) (*responses.ProxmoxResponse, error) {
		return session.PostForm(endpoint, form)
	})
}

func (proxmoxClient ProxmoxClient) PutForm(endpoint endpoints.Endpoint, form url.Values) (*responses.ProxmoxResponse, error) {
	return proxmoxClient.execute(func(session *ProxmoxSession) (*responses.ProxmoxResponse, error) {
		return session.PutForm(endpoint, form)
	})
}

func (proxmoxClient ProxmoxClient) Get(endpoint endpoints.Endpoint) (*responses.ProxmoxResponse, error) {
	return proxmoxClient.execute(func(session *ProxmoxSession) (*responses.ProxmoxResponse, error) {
		return session.Get(endpoint)
	})
}

func (proxmoxClient ProxmoxClient) execute(data func(session *ProxmoxSession) (*responses.ProxmoxResponse, error)) (*responses.ProxmoxResponse, error) {
	var session *ProxmoxSession
	var err error

	if proxmoxClient.currentSessionID == -1 {
		session, err = proxmoxClient.getOnlineSession()

		if err != nil {
			return nil, err
		}
	} else {
		session = proxmoxClient.sessions[proxmoxClient.currentSessionID]
	}

	response, err := data(session)

	if err != nil {
		proxError, ok := err.(error2.SessionOfflineError)

		if !ok {
			return nil, err
		}

		log.Println(proxError.Error())
		session, err = proxmoxClient.getOnlineSession()

		if err != nil {
			return nil, err
		}

		response, err = data(session)

		if err != nil {
			_, ok := err.(error2.SessionOfflineError)

			if !ok {
				return nil, err
			}

			return nil, error2.NoOnlineSessionError{SizeOfClients: len(proxmoxClient.sessions)}
		}
	}

	return response, nil
}
