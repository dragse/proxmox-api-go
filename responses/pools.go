package responses

type Pool struct {
	PoolID  string `json:"poolid"`
	Comment string `json:"comment"`
}

type PoolDetail struct {
	Members []*PoolMember `json:"members"`
	Comment string        `json:"comment"`
}

type PoolMember struct {
	ID        string  `json:"id"`
	VmID      int     `json:"vmid"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Node      string  `json:"node"`
	Status    string  `json:"status"`
	Uptime    int64   `json:"uptime"`
	Cpu       float64 `json:"cpu"`
	MaxCPU    int     `json:"maxcpu"`
	Template  int     `json:"template"`
	NetIn     int64   `json:"netin"`
	NetOut    int64   `json:"netout"`
	Mem       int64   `json:"mem"`
	MaxMem    int64   `json:"maxmem"`
	Disk      int64   `json:"disk"`
	MaxDisk   int64   `json:"maxdisk"`
	DiskRead  int64   `json:"diskread"`
	DiskWrite int64   `json:"diskwrite"`
}
