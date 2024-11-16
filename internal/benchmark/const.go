package benchmark

import "time"

const (
	// For both UDP and TCP.
	packetSize = 1024

	timeoutDurationUDP = 2 * time.Second
)
