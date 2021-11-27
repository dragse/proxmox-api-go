package ipset

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/client"
	error2 "github.com/dragse/proxmox-api-go/error"
	"github.com/dragse/proxmox-api-go/responses/node/vm"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"net/url"
	"regexp"
	"strconv"
)

type ProxmoxIPSet struct {
	client *client.ProxmoxClient

	NodeName  string
	VmID      int
	IPSetName string
}

func NewProxmoxIPSet(client *client.ProxmoxClient, nodeName string, vmID int, ipSetName string) *ProxmoxIPSet {
	return &ProxmoxIPSet{
		client:    client,
		NodeName:  nodeName,
		VmID:      vmID,
		IPSetName: ipSetName,
	}
}

func (ipSet ProxmoxIPSet) GetIPAddresses() ([]*vm.IPAddress, error) {
	var ipAddrs []*vm.IPAddress
	resp, err := ipSet.client.Get(endpoints.Nodes_Node_Qemu_VMID_FirewallIpset_Name_.FormatValues(ipSet.NodeName, strconv.Itoa(ipSet.VmID), ipSet.IPSetName))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &ipAddrs)

	if err != nil {
		return nil, err
	}

	return ipAddrs, nil
}

func (ipSet ProxmoxIPSet) Delete() error {
	_, err := ipSet.client.Delete(endpoints.Nodes_Node_Qemu_VMID_FirewallIpset_Name_.FormatValues(ipSet.NodeName, strconv.Itoa(ipSet.VmID), ipSet.IPSetName))

	if err != nil {
		return err
	}

	return nil
}

func (ipSet ProxmoxIPSet) AddIPNetwork(cidr string, comment string) error {
	matched, err := regexp.MatchString("\\b(([2]([0-4][0-9]|[5][0-5])|[0-1]?[0-9]?[0-9])[.]){3}(([2]([0-4][0-9]|[5][0-5])|[0-1]?[0-9]?[0-9]))\\b\\/\\b([0-9]|[12][0-9]|3[0-2])\\b", cidr)

	if err != nil {
		return err
	}

	if !matched {
		return error2.InvalidParameterError{Parameter: "cidr"}
	}

	params := url.Values{}
	params.Add("cidr", cidr)
	params.Add("comment", comment)

	_, err = ipSet.client.PostForm(endpoints.Nodes_Node_Qemu_VMID_FirewallIpset_Name_.FormatValues(ipSet.NodeName, strconv.Itoa(ipSet.VmID), ipSet.IPSetName), params)

	if err != nil {
		return err
	}

	return nil
}

func (ipSet ProxmoxIPSet) DeleteIPNetwork(cidr string) error {
	matched, err := regexp.MatchString("\\b(([2]([0-4][0-9]|[5][0-5])|[0-1]?[0-9]?[0-9])[.]){3}(([2]([0-4][0-9]|[5][0-5])|[0-1]?[0-9]?[0-9]))\\b(\\/\\b([0-9]|[12][0-9]|3[0-2]))?\\b", cidr)

	if err != nil {
		return err
	}

	if !matched {
		return error2.InvalidParameterError{Parameter: "cidr"}
	}

	_, err = ipSet.client.Delete(endpoints.Nodes_Node_Qemu_VMID_FirewallIpset_Name__Cidr_.FormatValues(ipSet.NodeName, strconv.Itoa(ipSet.VmID), ipSet.IPSetName, cidr))

	if err != nil {
		return err
	}

	return nil
}
