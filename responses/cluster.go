package responses

type ClusterStatusInformation struct {
	Name    string `json:"name"`
	Nodes   int    `json:"nodes"`
	Quorate int    `json:"quorate"`
	Version int    `json:"version"`
}
