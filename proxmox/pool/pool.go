package pool

import (
	"encoding/json"
	"github.com/dragse/proxmox-api-go/client"
	"github.com/dragse/proxmox-api-go/responses"
	"github.com/dragse/proxmox-api-go/static/endpoints"
)

type ProxmoxPool struct {
	client *client.ProxmoxClient

	PoolName string
}

func NewProxmoxPool(client *client.ProxmoxClient, name string) *ProxmoxPool {
	return &ProxmoxPool{client: client, PoolName: name}
}

func (pool ProxmoxPool) GetDetail() (*responses.PoolDetail, error) {
	var data *responses.PoolDetail
	resp, err := pool.client.Get(endpoints.Pools_Pool_.FormatValues(pool.PoolName))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
