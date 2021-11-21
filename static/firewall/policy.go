package firewall

type Policy string

const (
	Accept Policy = "ACCEPT"
	Reject Policy = "REJECT"
	Drop   Policy = "DROP"
)
