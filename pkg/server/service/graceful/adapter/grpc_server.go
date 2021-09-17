package adapter

import (
	"fmt"
	"log"
	"net"

	"github.com/ylallemant/panopticon/pkg/server/service/graceful"
	"google.golang.org/grpc"
)

func GRPCServer(server *grpc.Server, addr string) *grpcServerWrapper {
	return &grpcServerWrapper{
		server: server,
		addr:   addr,
	}
}

const grpcServerProcessName = "GRPC Server on %s"

var _ graceful.Process = &grpcServerWrapper{}

type grpcServerWrapper struct {
	server *grpc.Server
	addr   string
}

func (w *grpcServerWrapper) Name() string {
	return fmt.Sprintf(grpcServerProcessName, w.addr)
}

func (w *grpcServerWrapper) Start() error {
	go func() {
		lis, err := net.Listen("tcp", w.addr)
		if err != nil {
			log.Fatalf("failed to create lister for %s", w.Name())
		}

		if err := w.server.Serve(lis); err != nil {
			log.Printf("failed to serve %s", w.Name())
		}
	}()

	return nil
}

func (w *grpcServerWrapper) Stop() {
	w.server.GracefulStop()
}
