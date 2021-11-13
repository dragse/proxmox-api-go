package status

type Status string

const (
	Reboot   = "reboot"
	Reset    = "reset"
	Resume   = "resume"
	Shutdown = "shutdown"
	Start    = "start"
	Stop     = "stop"
	Suspend  = "suspend"
)
