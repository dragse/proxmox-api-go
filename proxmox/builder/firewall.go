package builder

import (
	"github.com/dragse/proxmox-api-go/static/firewall"
	"net/url"
	"strconv"
)

type FirewallBuilder struct {
	Enable      int
	DHCP        int
	NDP         int
	MacFilter   int
	IPFilter    int
	Radv        int
	LogLevelIn  firewall.LogLevel
	LogLevelOut firewall.LogLevel
	PolicyIn    firewall.Policy
	PolicyOut   firewall.Policy
}

func NewFirewallBuilder() *FirewallBuilder {
	return &FirewallBuilder{
		Enable:      0,
		DHCP:        0,
		NDP:         0,
		MacFilter:   0,
		IPFilter:    0,
		Radv:        0,
		LogLevelIn:  firewall.NoLog,
		LogLevelOut: firewall.NoLog,
		PolicyIn:    firewall.Drop,
		PolicyOut:   firewall.Accept,
	}
}

func (b *FirewallBuilder) SetEnable(enable bool) *FirewallBuilder {
	if enable {
		b.Enable = 1
	} else {
		b.Enable = 0
	}
	return b
}

func (b *FirewallBuilder) SetDHCP(dhcp bool) *FirewallBuilder {
	if dhcp {
		b.DHCP = 1
	} else {
		b.DHCP = 0
	}
	return b
}

func (b *FirewallBuilder) SetNDP(ndp bool) *FirewallBuilder {
	if ndp {
		b.NDP = 1
	} else {
		b.NDP = 0
	}
	return b
}

func (b *FirewallBuilder) SetMacFilter(macFilter bool) *FirewallBuilder {
	if macFilter {
		b.MacFilter = 1
	} else {
		b.MacFilter = 0
	}
	return b
}

func (b *FirewallBuilder) SetIPFilter(ipFilter bool) *FirewallBuilder {
	if ipFilter {
		b.IPFilter = 1
	} else {
		b.IPFilter = 0
	}
	return b
}

func (b *FirewallBuilder) SetRadv(radv bool) *FirewallBuilder {
	if radv {
		b.Radv = 1
	} else {
		b.Radv = 0
	}
	return b
}

func (b *FirewallBuilder) SetInputPolicy(inputPolicy firewall.Policy) *FirewallBuilder {
	b.PolicyIn = inputPolicy
	return b
}

func (b *FirewallBuilder) SetOutputPolicy(outputPolicy firewall.Policy) *FirewallBuilder {
	b.PolicyOut = outputPolicy
	return b
}

func (b *FirewallBuilder) SetGlobalLogLevel(logLevel firewall.LogLevel) *FirewallBuilder {
	b.LogLevelIn = logLevel
	b.LogLevelOut = logLevel
	return b
}

func (b *FirewallBuilder) SetInputLogLevel(logLevel firewall.LogLevel) *FirewallBuilder {
	b.LogLevelIn = logLevel
	return b
}

func (b *FirewallBuilder) SetOutputLogLevel(logLevel firewall.LogLevel) *FirewallBuilder {
	b.LogLevelOut = logLevel
	return b
}

func (b FirewallBuilder) BuildToValues() url.Values {
	params := url.Values{}
	params.Add("enable", strconv.Itoa(b.Enable))
	params.Add("dhcp", strconv.Itoa(b.DHCP))
	params.Add("ndp", strconv.Itoa(b.NDP))
	params.Add("macfilter", strconv.Itoa(b.MacFilter))
	params.Add("ipfilter", strconv.Itoa(b.IPFilter))
	params.Add("macfilter", strconv.Itoa(b.MacFilter))
	params.Add("radv", strconv.Itoa(b.Radv))
	params.Add("log_level_in", string(b.LogLevelIn))
	params.Add("log_level_out", string(b.LogLevelOut))
	params.Add("policy_in", string(b.PolicyIn))
	params.Add("policy_out", string(b.PolicyOut))
	return params
}
