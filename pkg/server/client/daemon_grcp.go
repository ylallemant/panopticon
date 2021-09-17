package client

import (
	"context"
	"log"
	"strings"
	"time"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"google.golang.org/grpc"
)

func NewGrcpDaemonClient(endpoint string, options *ClientOptions) (DaemonClient, error) {
	log.Printf("initialize gRPC daemon client for endpoint: %s", endpoint)
	client := new(grcpDaemonClient)
	client.endpoint = endpoint

	metadata, err := NewDaemonMetadata()
	if err != nil {
		return nil, err
	}
	client.metadata = metadata

	host := strings.Replace(endpoint, "grpc://", "", -1)
	log.Printf("target host: %s", host)

	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		return nil, err
	}

	client.conn = conn
	client.panopticon = v1.NewPanopticonServiceClient(client.conn)

	return client, nil
}

var _ DaemonClient = &grcpDaemonClient{}

type grcpDaemonClient struct {
	endpoint   string
	metadata   *DaemonMetadata
	panopticon v1.PanopticonServiceClient
	conn       *grpc.ClientConn
}

func (c *grcpDaemonClient) Type() string {
	return "grpc"
}

func (c *grcpDaemonClient) Metadata() *DaemonMetadata {
	return c.metadata
}

func (c *grcpDaemonClient) Report(report *v1.HostProcessReportRequest) (*v1.HostActionResponse, error) {
	log.Printf("sending report %s", time.Now())
	return c.panopticon.Report(context.Background(), report, grpc.EmptyCallOption{})
}

func (c *grcpDaemonClient) Stop() {
	c.conn.Close()
}
