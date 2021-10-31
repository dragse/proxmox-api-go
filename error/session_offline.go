package error

import "fmt"

type SessionOfflineError struct {
	Host string
}

func (s SessionOfflineError) Error() string {
	return fmt.Sprintf("session to host '%s' is offline", s.Host)
}
