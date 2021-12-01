package builder

import (
	"net/url"
	"strconv"
	"strings"
)

type ipConfig struct {
	gateway4 string
	ip4      string
	gateway6 string
	ip6      string
}

type CloudInitBuilder struct {
	password   string
	user       string
	ipConfigs  []*ipConfig
	nameserver string
	sshKeys    string
}

func NewCloudInitBuilder() *CloudInitBuilder {
	return &CloudInitBuilder{
		ipConfigs: make([]*ipConfig, 0),
	}
}

func (b *CloudInitBuilder) SetPassword(password string) *CloudInitBuilder {
	b.password = password
	return b
}

func (b *CloudInitBuilder) SetUser(user string) *CloudInitBuilder {
	b.user = user
	return b
}

func (b *CloudInitBuilder) AddIPConfig(ipv4 string, ipv4Gw string, ipv6 string, ipv6Gw string) *CloudInitBuilder {
	b.ipConfigs = append(b.ipConfigs, &ipConfig{
		gateway4: ipv4Gw,
		ip4:      ipv4,
		gateway6: ipv6Gw,
		ip6:      ipv6,
	})
	return b
}

func (b *CloudInitBuilder) SetNameservers(nameserver string) *CloudInitBuilder {
	b.nameserver = nameserver
	return b
}

func (b *CloudInitBuilder) AddSSHKey(sshKey string) *CloudInitBuilder {
	b.sshKeys = sshKey
	return b
}

func (b CloudInitBuilder) BuildToValues() url.Values {
	params := url.Values{}

	if b.password != "" {
		params.Add("cipassword", b.password)
	}

	if b.user != "" {
		params.Add("ciuser", b.user)
	}

	for i, config := range b.ipConfigs {
		args := make([]string, 0)

		if config.ip4 != "" {
			args = append(args, "ip="+config.ip4)
		}

		if config.ip6 != "" {
			args = append(args, "ip6="+config.ip6)
		}

		if config.gateway4 != "" {
			args = append(args, "gw="+config.gateway4)
		}

		if config.gateway6 != "" {
			args = append(args, "gw6="+config.gateway6)
		}

		if len(args) == 0 {
			continue
		}

		params.Add("ipconfig"+strconv.Itoa(i), strings.Join(args, ","))
	}

	if b.nameserver != "" {
		params.Add("nameserver", b.nameserver)
	}

	if b.sshKeys != "" {
		params.Add("sshkeys", b.sshKeys)
	}

	return params
}
