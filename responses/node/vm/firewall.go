package vm

import "github.com/dragse/proxmox-api-go/static/firewall"

type FirewallLog struct {
	N int    `json:"n"`
	T string `json:"t"`
}

type FirewallOption struct {
	Digest      string            `json:"digest"`
	Enable      int               `json:"enable"`
	DHCP        int               `json:"dhcp"`
	NDP         int               `json:"ndp"`
	MacFilter   int               `json:"macfilter"`
	IPFilter    int               `json:"ipfilter"`
	Radv        int               `json:"radv"`
	LogLevelIn  firewall.LogLevel `json:"log_level_in"`
	LogLevelOut firewall.LogLevel `json:"log_level_out"`
	PolicyIn    firewall.Policy   `json:"policy_in"`
	PolicyOut   firewall.Policy   `json:"policy_out"`
}
