package endpoints

import "regexp"

type Endpoint string

var /* const */ placeholderRegex = regexp.MustCompile("{val}")

func (endpoint Endpoint) FormatValues(val ...string) Endpoint {
	i := 0
	format := placeholderRegex.ReplaceAllStringFunc(string(endpoint), func(s string) string {
		return val[i]
	})

	return Endpoint(format)
}
