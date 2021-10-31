package responses

type NodeStatusInformation struct {
	IP     string `json:"ip"`
	Level  string `json:"level"`
	Local  int    `json:"local"`
	Name   string `json:"name"`
	NodeID int    `json:"node_id"`
	Online int    `json:"online"`
}

type NodeInformation struct {
	Node           string  `json:"node"`
	Status         string  `json:"status"` // unknown / online / offline
	CPU            float64 `json:"cpu"`
	Level          string  `json:"level"`
	MaxCPU         int     `json:"maxcpu"`
	Mem            int64   `json:"mem"`     // used in bytes
	MaxMem         int64   `json:"maxmem"`  // in bytes
	Disk           int64   `json:"disk"`    // used in bytes
	MaxDisk        int64   `json:"maxdisk"` // used in bytes
	SSLFingerprint string  `json:"ssl_fingerprint"`
	Uptime         int64   `json:"uptime"` // in seconds
}
