package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/dragse/proxmox-api-go/responses"
	"github.com/dragse/proxmox-api-go/static"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ProxmoxSession struct {
	Hostname string
	Username string
	Token string

	VerifySSL bool
	Client *http.Client
}

func (proxmoxHost *ProxmoxSession) SetupClient() error  {
	var tr *http.Transport

	if proxmoxHost.VerifySSL {
		tr = &http.Transport{
			DisableKeepAlives:      false,
			IdleConnTimeout:        0,
			MaxIdleConns:           200,
			MaxIdleConnsPerHost:    100,
		}
	} else {
		tr = &http.Transport{
			DisableKeepAlives:      false,
			IdleConnTimeout:        0,
			MaxIdleConns:           200,
			MaxIdleConnsPerHost:    100,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	proxmoxHost.Client = &http.Client{
		Transport:     tr,
	}

	return nil
}

func (proxmox ProxmoxSession) TestConnection() error {
	_, err := proxmox.Get(static.EndpointVersion)

	if err != nil {
		return err
	}

	return nil
}

func (host ProxmoxSession) PostForm(endpoint static.Endpoint, form url.Values) (*responses.ProxmoxResponse, error) {
	var target string
	var data responses.ProxmoxResponse
	var req *http.Request

	target = "https://" + host.Hostname + "/api2/json" + string(endpoint)

	req, err := http.NewRequest("POST", target, bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	req.Header.Add("Authorization", "PVEAPIToken=" + host.Username + "=" + host.Token)

	r, err := host.Client.Do(req)

	if err != nil {
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

func (host ProxmoxSession) Get(endpoint static.Endpoint) (*responses.ProxmoxResponse, error) {
	var target string
	var data responses.ProxmoxResponse

	target = "https://" + host.Hostname + "/api2/json" + string(endpoint)

	req, err := http.NewRequest("GET", target, nil)

	req.Header.Add("Authorization", "PVEAPIToken=" + host.Username + "=" + host.Token)

	r, err := host.Client.Do(req)

	if err != nil {
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