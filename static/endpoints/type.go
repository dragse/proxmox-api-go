package endpoints

import "regexp"

type Endpoint string

var /* const */ placeholderRegex = regexp.MustCompile("{val}")

func (endpoint Endpoint) FormatValues(val ...string) Endpoint {
	i := -1
	format := placeholderRegex.ReplaceAllStringFunc(string(endpoint), func(s string) string {
		i += 1
		return val[i]
	})

	return Endpoint(format)
}
