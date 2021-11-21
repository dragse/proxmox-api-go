package firewall

type LogLevel string

const (
	Alert    LogLevel = "alert"
	Critical LogLevel = "crit"
	Error    LogLevel = "err"
	Warning  LogLevel = "warning"
	Notice   LogLevel = "notice"
	Info     LogLevel = "info"
	Debug    LogLevel = "debug"
	NoLog    LogLevel = "nolog"
)
