package vm

type VNCProxy struct {
	Cert   string `json:"cert"`
	Port   string `json:"port"`
	Ticket string `json:"ticket"`
	UpID   string `json:"upid"`
	User   string `json:"user"`
}
