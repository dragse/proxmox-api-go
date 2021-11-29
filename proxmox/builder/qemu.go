package builder

import (
	"fmt"
	"github.com/dragse/proxmox-api-go/static/operation_system"
	"github.com/dragse/proxmox-api-go/util"
	"net/url"
	"strconv"
)

type storage struct {
	storageType string
	storageSize string
}

type VmBuilder struct {
	id       string
	name     string
	osType   operation_system.OSType
	cpuType  string
	cores    int
	sockets  int
	memory   *util.Byte
	networks []string
	storages []*storage
	iso      *storage
	pool     string
}

func NewVmBuilder() *VmBuilder {
	return &VmBuilder{
		osType:   operation_system.Other,
		cores:    1,
		sockets:  1,
		networks: make([]string, 0),
		memory:   util.NewBytesFromMegaBytes(512),
	}
}

func (b *VmBuilder) SetID(id string) *VmBuilder {
	b.id = id
	return b
}

func (b *VmBuilder) SetName(name string) *VmBuilder {
	b.name = name
	return b
}

func (b *VmBuilder) SetCPUType(cpu string) *VmBuilder {
	b.cpuType = cpu
	return b
}

func (b *VmBuilder) SetCoresPerSocket(cores int) *VmBuilder {
	b.cores = cores
	return b
}

func (b *VmBuilder) SetSocket(cores int) *VmBuilder {
	b.sockets = cores
	return b
}

func (b *VmBuilder) SetOSType(osType operation_system.OSType) *VmBuilder {
	b.osType = osType
	return b
}

func (b *VmBuilder) SetMemory(memory *util.Byte) *VmBuilder {
	b.memory = memory
	return b
}

func (b *VmBuilder) SetPool(pool string) *VmBuilder {
	b.pool = pool
	return b
}

func (b *VmBuilder) SetIso(disk string, isoFile string) *VmBuilder {
	b.iso = &storage{storageType: disk, storageSize: isoFile}
	return b
}

func (b *VmBuilder) AddNetwork(networkBridge string) *VmBuilder {
	b.networks = append(b.networks, networkBridge)
	return b
}

func (b *VmBuilder) AddStorage(disk string, size string) *VmBuilder {
	b.storages = append(b.storages, &storage{storageType: disk, storageSize: size})
	return b
}

func (b VmBuilder) BuildToValues() url.Values {
	params := url.Values{}

	if b.id != "" {
		params.Add("vmid", b.id)
	}

	params.Add("cores", strconv.Itoa(b.cores))
	params.Add("ostype", string(b.osType))
	params.Add("scsihw", "virtio-scsi-pci")
	params.Add("sockets", strconv.Itoa(b.sockets))
	params.Add("memory", strconv.FormatInt(b.memory.ToMegaByte(), 10))

	if b.cpuType != "" {
		params.Add("cpu", b.cpuType)
	}

	if b.name != "" {
		params.Add("name", b.name)
	}

	if b.pool != "" {
		params.Add("pool", b.pool)
	}

	if b.iso != nil {
		params.Add("ide2", b.iso.storageType+":iso/"+b.iso.storageSize+",media=cdrom")
	}

	for i, storage := range b.storages {
		params.Add("scsi"+strconv.Itoa(i), fmt.Sprintf("%s:%s", storage.storageType, storage.storageSize))
	}

	for i, network := range b.networks {
		params.Add("net"+strconv.Itoa(i), "virtio,bridge="+network+",firewall=1")
	}

	return params
}
