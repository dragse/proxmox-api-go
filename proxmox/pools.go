package proxmox

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/responses"
	"github.com/dragse/proxmox-api-go/static/endpoints"
	"net/url"
)

func (proxmoxCluster ProxmoxCluster) CreatePool(poolName string, comment string) error {
	form := url.Values{
		"poolid":  {poolName},
		"comment": {comment},
	}

	_, err := proxmoxCluster.client.PostForm(endpoints.Pools, form)

	if err != nil {
		return err
	}

	return nil
}

func (proxmoxCluster ProxmoxCluster) GetPools() ([]*responses.Pool, error) {
	var pools []*responses.Pool
	resp, err := proxmoxCluster.client.Get(endpoints.Pools)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &pools)

	if err != nil {
		return nil, err
	}

	return pools, nil
}
