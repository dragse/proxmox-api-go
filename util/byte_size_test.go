package util

import (
	"reflect"
	"testing"
)

func TestByte_ToBytes(t *testing.T) {
	type fields struct {
		bytes int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name:   "zero bytes",
			fields: fields{bytes: 0},
			want:   0,
		},
		{
			name:   "one kb",
			fields: fields{bytes: 1000},
			want:   1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Byte{
				bytes: tt.fields.bytes,
			}
			if got := b.ToBytes(); got != tt.want {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByte_ToGigaByte(t *testing.T) {
	type fields struct {
		bytes int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name:   "zero bytes",
			fields: fields{bytes: 0},
			want:   0,
		},
		{
			name:   "one gb",
			fields: fields{bytes: 1000 * 1000 * 1000},
			want:   1,
		},
		{
			name:   "one tb",
			fields: fields{bytes: 1000 * 1000 * 1000 * 1000},
			want:   1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Byte{
				bytes: tt.fields.bytes,
			}
			if got := b.ToGigaByte(); got != tt.want {
				t.Errorf("ToGigaByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByte_ToMegaByte(t *testing.T) {
	type fields struct {
		bytes int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name:   "zero bytes",
			fields: fields{bytes: 0},
			want:   0,
		},
		{
			name:   "one mb",
			fields: fields{bytes: 1000 * 1000},
			want:   1,
		},
		{
			name:   "one gb",
			fields: fields{bytes: 1000 * 1000 * 1000},
			want:   1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Byte{
				bytes: tt.fields.bytes,
			}
			if got := b.ToMegaByte(); got != tt.want {
				t.Errorf("ToMegaByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBytesFromBytes(t *testing.T) {
	type args struct {
		bytes int64
	}
	tests := []struct {
		name string
		args args
		want *Byte
	}{
		{
			name: "zero bytes",
			args: args{bytes: 0},
			want: &Byte{bytes: 0},
		},
		{
			name: "one kb",
			args: args{bytes: 1000},
			want: &Byte{bytes: 1000},
		},
		{
			name: "one mb",
			args: args{bytes: 1000 * 1000},
			want: &Byte{bytes: 1000 * 1000},
		},
		{
			name: "one gb",
			args: args{bytes: 1000 * 1000 * 1000},
			want: &Byte{bytes: 1000 * 1000 * 1000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBytesFromBytes(tt.args.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBytesFromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBytesFromGigaBytes(t *testing.T) {
	type args struct {
		bytes int64
	}
	tests := []struct {
		name string
		args args
		want *Byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBytesFromGigaBytes(tt.args.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBytesFromGigaBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBytesFromMegaBytes(t *testing.T) {
	type args struct {
		bytes int64
	}
	tests := []struct {
		name string
		args args
		want *Byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBytesFromMegaBytes(tt.args.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBytesFromMegaBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
