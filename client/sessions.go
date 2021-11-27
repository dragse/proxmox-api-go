package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	error2 "github.com/dragse/proxmox-api-go/error"
	"github.com/dragse/proxmox-api-go/responses"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ProxmoxSession struct {
	Hostname string
	Username string
	Token    string

	VerifySSL bool
	Client    *http.Client
}

func (proxmoxHost *ProxmoxSession) SetupClient() error {
	var tr *http.Transport

	if proxmoxHost.VerifySSL {
		tr = &http.Transport{
			DisableKeepAlives:   false,
			IdleConnTimeout:     0,
			MaxIdleConns:        200,
			MaxIdleConnsPerHost: 100,
		}
	} else {
		tr = &http.Transport{
			DisableKeepAlives:   false,
			IdleConnTimeout:     0,
			MaxIdleConns:        200,
			MaxIdleConnsPerHost: 100,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		}
	}

	proxmoxHost.Client = &http.Client{
		Transport: tr,
	}

	return nil
}

func (proxmox ProxmoxSession) TestConnection() error {
	_, err := proxmox.Get(endpoints.EndpointVersion)

	if err != nil {
		return err
	}

	return nil
}

func (host ProxmoxSession) PostForm(endpoint endpoints.Endpoint, form url.Values) (*responses.ProxmoxResponse, error) {
	var target string
	var req *http.Request

	target = host.formatProxmoxAPI(endpoint)

	req, err := http.NewRequest("POST", target, bytes.NewBufferString(form.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	req.Header.Add("Authorization", "PVEAPIToken="+host.Username+"="+host.Token)

	return host.handleRequest(req)
}

func (host ProxmoxSession) PutForm(endpoint endpoints.Endpoint, form url.Values) (*responses.ProxmoxResponse, error) {
	var target string
	var req *http.Request

	target = host.formatProxmoxAPI(endpoint)

	req, err := http.NewRequest("PUT", target, bytes.NewBufferString(form.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	req.Header.Add("Authorization", "PVEAPIToken="+host.Username+"="+host.Token)

	return host.handleRequest(req)
}

func (host ProxmoxSession) Get(endpoint endpoints.Endpoint) (*responses.ProxmoxResponse, error) {
	target := host.formatProxmoxAPI(endpoint)

	req, err := http.NewRequest("GET", target, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "PVEAPIToken="+host.Username+"="+host.Token)

	return host.handleRequest(req)
}

func (host ProxmoxSession) Delete(endpoint endpoints.Endpoint) (*responses.ProxmoxResponse, error) {
	target := host.formatProxmoxAPI(endpoint)

	req, err := http.NewRequest("DELETE", target, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "PVEAPIToken="+host.Username+"="+host.Token)

	return host.handleRequest(req)
}

func (host ProxmoxSession) formatProxmoxAPI(endpoint endpoints.Endpoint) string {
	return "https://" + host.Hostname + "/api2/json" + string(endpoint)
}

func (host ProxmoxSession) handleRequest(request *http.Request) (*responses.ProxmoxResponse, error) {
	var data responses.ProxmoxResponse

	r, err := host.Client.Do(request)

	if err != nil {
		urlError := err.(*url.Error)

		if urlError.Timeout() {
			return nil, error2.SessionOfflineError{
				Host: host.Hostname,
			}
		}

		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, errors.New("HTTP Error " + r.Status)
	}

	response, err := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
