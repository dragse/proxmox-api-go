package util

type Byte struct {
	bytes int64
}

func NewBytesFromBytes(bytes int64) *Byte {
	return &Byte{bytes: bytes}
}

func NewBytesFromMegaBytes(bytes int64) *Byte {
	return &Byte{bytes: bytes * 1000}
}

func NewBytesFromGigaBytes(bytes int64) *Byte {
	return &Byte{bytes: bytes * 1000 * 1000}
}

func (b Byte) ToBytes() int64 {
	return b.bytes
}

func (b Byte) ToMegaByte() int64 {
	return b.bytes / 1000
}

func (b Byte) ToGigaByte() int64 {
	return b.ToMegaByte() / 1000
}
