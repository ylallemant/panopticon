package client

import (
	"os"
	"runtime"
)

func NewDaemonMetadata() (*DaemonMetadata, error) {
	meta := new(DaemonMetadata)

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	meta.hostname = hostname
	meta.arch = runtime.GOARCH
	meta.os = runtime.GOOS

	return meta, nil
}

type DaemonMetadata struct {
	hostname string
	host     string
	arch     string
	os       string
}

func (m *DaemonMetadata) Hostname() string {
	return m.hostname
}

func (m *DaemonMetadata) Host() string {
	return m.host
}

func (m *DaemonMetadata) Arch() string {
	return m.os
}

func (m *DaemonMetadata) OS() string {
	return m.os
}
