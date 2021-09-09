package api

import "time"

type Process struct {
	PID       int
	PPID      int
	UserID    int
	User      string
	StartTime time.Time
	Command   string
}

func (s Process) Pid() int {
	return s.PID
}

func (s Process) PPid() int {
	return s.PID
}

func (s Process) Executable() string {
	return s.Command
}
