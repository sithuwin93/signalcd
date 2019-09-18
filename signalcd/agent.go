package signalcd

import (
	"time"
)

// Agent is a daemon responsible for running the pipelines in namespaces
type Agent struct {
	Name      string
	Heartbeat time.Time
}

// Ready returns if an Agent is ready.
// If ready an agent has reported to the API in the past 15s.
func (a Agent) Ready() bool {
	return time.Since(a.Heartbeat) < 15*time.Second
}
