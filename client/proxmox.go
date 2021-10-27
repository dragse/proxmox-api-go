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
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

type ProxmoxHost struct {
	Hostname string
	Username string
	Password string

	VerifySSL bool
	Client *http.Client

	ticket string
	csrfPreventionToken string
}

func (proxmoxHost *ProxmoxHost) Login() error  {
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

	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}

	proxmoxHost.Client = &http.Client{
		Transport:     tr,
		Jar: jar,
	}

	form := url.Values{
		"username": {proxmoxHost.Username},
		"password": {proxmoxHost.Password},
	}

	data, err := proxmoxHost.postForm(static.EndpointAccessTicket, form)
	if err != nil {
		return err
	}

	m := data.Data.(map[string]interface{})
	proxmoxHost.ticket = m["ticket"].(string)
	proxmoxHost.csrfPreventionToken = m["CSRFPreventionToken"].(string)

	cookie := &http.Cookie{
		Name:  "PVEAuthCookie",
		Value: proxmoxHost.ticket,
		Path:  "/",
	}

	cookieURL, err := url.Parse("https://" + proxmoxHost.Hostname + "/")
	if err != nil {
		return err
	}

	proxmoxHost.Client.Jar.SetCookies(cookieURL, []*http.Cookie{cookie})
	return nil
}

func (host ProxmoxHost) postForm(endpoint static.Endpoint, form url.Values) (*responses.ProxmoxResponse, error) {
	var target string
	var data responses.ProxmoxResponse
	var req *http.Request

	target = "https://" + host.Hostname + "/api2/json" + string(endpoint)

	req, err := http.NewRequest("POST", target, bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if host.csrfPreventionToken != "" {
		req.Header.Add("CSRFPreventionToken", host.csrfPreventionToken)
	}

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