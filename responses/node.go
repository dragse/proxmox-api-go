package responses

type NodeStatusInformation struct {
	IP     string `json:"ip"`
	Level  string `json:"level"`
	Local  int    `json:"local"`
	Name   string `json:"name"`
	NodeID int    `json:"node_id"`
	Online int    `json:"online"`
}
