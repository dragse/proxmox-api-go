package node

type CPUInfo struct {
	Sockets int    `json:"sockets"`
	UserHZ  int    `json:"user_hz"`
	Mhz     string `json:"mhz"`
	Cores   int    `json:"cores"`
	Cpus    int    `json:"cpus"`
	Hvm     string `json:"hvm"`
	Flags   string `json:"flags"`
	Model   string `json:"model"`
}

type ByteStatus struct {
	Free  int64 `json:"free"`
	Total int64 `json:"total"`
	Used  int64 `json:"used"`
	Avail int64 `json:"avail"`
}
