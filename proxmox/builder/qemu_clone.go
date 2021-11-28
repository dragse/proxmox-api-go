package builder

import (
	"net/url"
	"strconv"
)

type VmCopyBuilder struct {
	newId       int
	name        string
	description string
	full        bool
	pool        string
	target      string
}

func NewVmCopyBuilder() *VmCopyBuilder {
	return &VmCopyBuilder{}
}

func (b *VmCopyBuilder) SetNewID(id int) *VmCopyBuilder {
	b.newId = id
	return b
}

func (b *VmCopyBuilder) SetName(name string) *VmCopyBuilder {
	b.name = name
	return b
}

func (b *VmCopyBuilder) SetDescription(description string) *VmCopyBuilder {
	b.description = description
	return b
}

func (b *VmCopyBuilder) SetFullCopy(full bool) *VmCopyBuilder {
	b.full = full
	return b
}

func (b *VmCopyBuilder) SetPool(pool string) *VmCopyBuilder {
	b.pool = pool
	return b
}

func (b *VmCopyBuilder) SetTargetNode(node string) *VmCopyBuilder {
	b.target = node
	return b
}

func (b VmCopyBuilder) BuildToValues() url.Values {
	params := url.Values{}
	params.Add("newid", strconv.Itoa(b.newId))

	if b.full {
		params.Add("full", "1")
	} else {
		params.Add("full", "0")
	}

	if b.pool != "" {
		params.Add("pool", b.pool)
	}

	if b.name != "" {
		params.Add("name", b.name)
	}

	if b.description != "" {
		params.Add("description", b.description)
	}

	if b.target != "" {
		params.Add("target", b.target)
	}

	return params
}
