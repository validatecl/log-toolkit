package log_toolkit

import "os"

// OsInformation estructura con información de Sistema
type OsInformation struct {
	Pid      int
	Hostname string
}

// NewOsInformation Obtiene información de Hostname y Process ID
func NewOsInformation() *OsInformation {
	hostname, _ := os.Hostname()
	return &OsInformation{
		Pid:      os.Getpid(),
		Hostname: hostname,
	}
}
