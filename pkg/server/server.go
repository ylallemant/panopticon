package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/cli/server/options"
	"github.com/ylallemant/panopticon/pkg/server/analyse"
	"github.com/ylallemant/panopticon/pkg/server/cache"
	"github.com/ylallemant/panopticon/pkg/server/classifier"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func NewServer(options *options.Options) *server {
	process := new(server)

	process.ports.GRPC = options.Ports.GRPC
	process.ports.HTTP = options.Ports.HTTP
	process.ports.Maintenance = options.Ports.Maintenance

	process.cache = cache.NewMemoryCache()
	process.classifier = classifier.NewRegexpClassifier(process.cache)

	logger := logrus.NewEntry(logrus.New())
	grpc_logrus.ReplaceGrpcLogger(logger)
	process.grpcServer = grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_prometheus.StreamServerInterceptor,
			grpc_logrus.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		),
	)
	v1.RegisterPanopticonServiceServer(process.grpcServer, process)
	grpc_prometheus.Register(process.grpcServer)

	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	runtime.SetHTTPBodyMarshaler(gwmux)

	mux := http.NewServeMux()
	mux.Handle("/", wsproxy.WebsocketProxy(gwmux))

	process.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", process.ports.HTTP),
		Handler: mux,
	}

	process.promServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", process.ports.Maintenance),
		Handler: promhttp.Handler(),
	}

	return process
}

var _ v1.PanopticonServiceServer = &server{}

type server struct {
	ports      options.Ports
	errorGroup errgroup.Group
	grpcServer *grpc.Server
	httpServer *http.Server
	promServer *http.Server
	cache      cache.Cache
	classifier classifier.Classifier
	v1.UnimplementedPanopticonServiceServer
	stopCh chan struct{}
}

func (s *server) Name() string {
	return "panopticon-service"
}

func (s *server) Start() error {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.ports.GRPC))
	if err != nil {
		return err
	}

	s.errorGroup.Go(func() error {
		log.Printf("%s listening on gRPC port %d", s.Name(), s.ports.GRPC)
		return s.grpcServer.Serve(lis)
	})
	s.errorGroup.Go(func() error {
		log.Printf("%s listening on HTTP port %d", s.Name(), s.ports.HTTP)
		return s.httpServer.ListenAndServe()
	})
	s.errorGroup.Go(func() error {
		log.Printf("%s listening on maintenance port %d", s.Name(), s.ports.Maintenance)
		return s.promServer.ListenAndServe()
	})

	return s.errorGroup.Wait()
}

func (s *server) Stop() {
	ctx := context.Background()

	s.grpcServer.GracefulStop()
	s.httpServer.Shutdown(ctx)
	s.promServer.Shutdown(ctx)
	s.stopCh <- struct{}{}
}

func (s *server) Report(ctx context.Context, report *v1.HostProcessReportRequest) (*v1.HostActionResponse, error) {
	response := new(v1.HostActionResponse)

	log.Print("------")
	log.Printf("hostname:  %s", report.Report.Hostname)
	log.Printf("Arch:      %s", report.Report.Arch)
	log.Printf("OS:        %s", report.Report.Os)
	log.Printf("Processes: %d", len(report.Report.Processes))
	err := analyse.Analyse(report, s.cache, s.classifier)

	return response, err
}
