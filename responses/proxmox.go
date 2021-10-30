package responses

import "encoding/json"

type ProxmoxResponse struct {
	Data *json.RawMessage `json:"data"`
}

type Version struct {
	Version string `json:"version"`
	Release string `json:"release"`
	Keyboard string `json:"keyboard"`
	RepoID string `json:"repo_id"`
}
