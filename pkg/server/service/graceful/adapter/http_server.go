package adapter

import (
	"net/http"

	"context"

	"net"

	"fmt"

	"strings"

	"log"
	"time"

	"github.com/ylallemant/panopticon/pkg/server/service/graceful"
)

const httpServerProcessName = "HTTP Server on %s"

var _ graceful.Process = &httpServer{}

func HTTPServer(s *http.Server, addr string) *httpServer {
	return &httpServer{addr, s}
}

type httpServer struct {
	addr   string
	server *http.Server
}

func (w *httpServer) Name() string {
	return fmt.Sprintf(httpServerProcessName, w.addr)
}

func (w *httpServer) Start() error {
	go func() {
		lis, err := net.Listen("tcp", w.addr)
		if err != nil {
			log.Fatalf("failed to create lister for %s", w.Name())
		}

		if err := w.server.Serve(lis); err != nil {
			if !strings.Contains(err.Error(), "Server closed") {
				log.Printf("%s failed to serve due to: %s", w.Name(), err)
			}
		}
	}()

	return nil
}

func (w *httpServer) Stop() {
	// wait a second before killing the socket to give kubernetes services time to
	// rewrite iptables (kube-proxy) and refresh active pods (ingresses)
	time.Sleep(time.Second)
	w.server.Shutdown(context.Background())
}
