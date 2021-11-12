package vm

type Information struct {
	VmID      int     `json:"vmid"`
	Name      string  `json:"name"`
	Cpus      int     `json:"cpus"`
	Cpu       float64 `json:"cpu"`
	Status    string  `json:"status"` // stopped | running
	Lock      string  `json:"lock"`
	MaxMem    int64   `json:"maxmem"`  // in bytes
	Mem       int64   `json:"mem"`     // In bytes
	MaxDisk   int64   `json:"maxdisk"` // In bytes
	Disk      int64   `json:"disk"`    // In bytes
	Pid       int     `json:"pid"`
	NetIn     int64   `json:"netin"`
	NetOut    int64   `json:"netout"`
	DiskRead  int64   `json:"diskread"`
	DiskWrite int64   `json:"diskwrite"`
	Uptime    int     `json:"uptime"` // in seconds
}

type Detail struct {
	VmID      int       `json:"vmid"`
	Name      string    `json:"name"`
	Cpus      int       `json:"cpus"`
	Cpu       float64   `json:"cpu"`
	Status    string    `json:"status"`  // stopped | running
	MaxMem    int64     `json:"maxmem"`  // in bytes
	Mem       int64     `json:"mem"`     // In bytes
	MaxDisk   int64     `json:"maxdisk"` // In bytes
	Disk      int64     `json:"disk"`    // In bytes
	Pid       int       `json:"pid"`
	NetIn     int64     `json:"netin"`
	NetOut    int64     `json:"netout"`
	DiskRead  int64     `json:"diskread"`
	DiskWrite int64     `json:"diskwrite"`
	Uptime    int       `json:"uptime"` // in seconds
	HA        *DetailHA `json:"ha"`
}

type DetailHA struct {
	State   string `json:"state"`
	Group   string `json:"group"`
	Managed int    `json:"managed"`
}
