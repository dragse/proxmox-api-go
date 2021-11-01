package node

type StatusInformation struct {
	IP     string `json:"ip"`
	Level  string `json:"level"`
	Local  int    `json:"local"`
	Name   string `json:"name"`
	NodeID int    `json:"node_id"`
	Online int    `json:"online"`
}

type Information struct {
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

type Detail struct {
	Idle       int         `json:"idle"`
	CPU        float64     `json:"cpu"`
	PVEVersion string      `json:"pveversion"`
	Uptime     int         `json:"uptime"`
	KVersion   string      `json:"kversion"`
	Wait       int         `json:"wait"`
	LoadAVG    []string    `json:"loadavg"`
	CPUInfo    *CPUInfo    `json:"cpu_info"`
	Memory     *ByteStatus `json:"memory"`
	Swap       *ByteStatus `json:"swap"`
	RootFS     *ByteStatus `json:"rootfs"`
}

type TimeInformation struct {
	Localtime int64  `json:"localtime"`
	Time      int64  `json:"time"`
	Timezone  string `json:"timezone"`
}

type Version struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
}
