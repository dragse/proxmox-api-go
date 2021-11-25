package endpoints

import "testing"

func TestEndpoint_FormatValues(t *testing.T) {
	type args struct {
		val []string
	}
	tests := []struct {
		name     string
		endpoint Endpoint
		args     args
		want     Endpoint
	}{
		{
			name:     "Empty Endpoint",
			endpoint: "/empty",
			args:     args{val: []string{}},
			want:     "/empty",
		},
		{
			name:     "Endpoint for one Argument",
			endpoint: "/one/{val}",
			args: args{
				[]string{"args"},
			},
			want: "/one/args",
		},
		{
			name:     "Endpoint with two Arguments",
			endpoint: "/one/{val}/and/{val}",
			args:     args{[]string{"arg", "another"}},
			want:     "/one/arg/and/another",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.endpoint.FormatValues(tt.args.val...); got != tt.want {
				t.Errorf("FormatValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
