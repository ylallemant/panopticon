package client

import (
	"log"
	"strings"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
)

func NewDaemonClient(endpoint string, options *ClientOptions) (DaemonClient, error) {
	log.Printf("initialize client for endpoint: %s", endpoint)
	if strings.HasPrefix(endpoint, "grpc://") {
		return NewGrcpDaemonClient(endpoint, options)
	} else {
		return NewHttpDaemonClient(endpoint, options)
	}
}

type DaemonClient interface {
	Type() string
	Metadata() *DaemonMetadata
	Report(*v1.HostProcessReportRequest) (*v1.HostActionResponse, error)
	Stop()
}

type ClientOptions struct {
	Endpoint string
	Unsecure bool
}
