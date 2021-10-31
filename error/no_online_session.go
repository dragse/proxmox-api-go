package error

import "fmt"

type NoOnlineSessionError struct {
	SizeOfClients int
}

func (err NoOnlineSessionError) Error() string {
	return fmt.Sprintf("no active session of cluster found. checked %d sessions", err.SizeOfClients)
}
