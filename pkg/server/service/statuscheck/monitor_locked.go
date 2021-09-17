package statuscheck

import (
	"sync"
	"time"

	"github.com/spf13/pflag"
)

var lockDuration = pflag.Duration("shutdown-lock-time", 5*time.Second, "Number of seconds to block shutdown procedure")

func NewLockingStatusMonitor(name string) *lockingStatusMonitor {
	return &lockingStatusMonitor{
		name:         name,
		mu:           sync.Mutex{},
		status:       Bad,
		lockDuration: *lockDuration,
	}
}

type Status string

type lockingStatusMonitor struct {
	name         string
	mu           sync.Mutex
	status       Status
	lockDuration time.Duration
}

func (s *lockingStatusMonitor) Name() string {
	return s.name
}

func (s *lockingStatusMonitor) Start() error {
	s.status = Good
	return nil
}

func (s *lockingStatusMonitor) Stop() {
	s.status = Bad
	<-time.After(s.lockDuration)
}

func (s *lockingStatusMonitor) Status() Status {
	return s.status
}
