package client

import (
	"log"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
)

func NewHttpDaemonClient(endpoint string, options *ClientOptions) (DaemonClient, error) {
	log.Printf("initialize HTTP daemon client for endpoint: %s", endpoint)
	client := new(httpDaemonClient)
	client.endpoint = endpoint

	metadata, err := NewDaemonMetadata()
	if err != nil {
		return nil, err
	}
	client.metadata = metadata

	return client, nil
}

var _ DaemonClient = &httpDaemonClient{}

type httpDaemonClient struct {
	endpoint string
	metadata *DaemonMetadata
}

func (c *httpDaemonClient) Type() string {
	return "http"
}

func (c *httpDaemonClient) Metadata() *DaemonMetadata {
	return c.metadata
}

func (c *httpDaemonClient) Report(*v1.HostProcessReportRequest) (*v1.HostActionResponse, error) {
	return &v1.HostActionResponse{}, nil
}

func (c *httpDaemonClient) Stop() {
}
