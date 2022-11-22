package grpc

import (
	"runtime"
	"time"
)

const CleanupInterval = time.Minute * 60

func (s *Service) cleanupConnections() {
	for {
		time.Sleep(CleanupInterval)
		runtime.GC()
	}
}
