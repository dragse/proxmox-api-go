package error

type ProxmoxError struct {
	Err string
}

func (prox ProxmoxError) Error() string {
	return prox.Err
}
