package daemon

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/panopticon/pkg/cli/daemon/options"
	serverOtions "github.com/ylallemant/panopticon/pkg/cli/server/options"
	runtime "github.com/ylallemant/panopticon/pkg/daemon"
	"github.com/ylallemant/panopticon/pkg/server"
	"github.com/ylallemant/panopticon/pkg/server/service/graceful"
)

var rootCmd = &cobra.Command{
	Use:   "daemon",
	Short: "used to monitor a host",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		startDefaultServer := false

		if options.Current.Endpoint == "" {
			options.Current.Endpoint = fmt.Sprintf("grpc://localhost:%d", serverOtions.Current.Ports.GRPC)
			startDefaultServer = true
		}

		processes := graceful.NewProcessBucket()

		deamon, err := runtime.NewDaemon(options.Current)
		if err != nil {
			return err
		}

		processes.AddProcess(deamon)

		if startDefaultServer {
			serverOptions := serverOtions.Current
			service := server.NewServer(serverOptions)
			processes.AddProcess(service)
		}

		signal.Notify(processes.SigChan(), syscall.SIGTERM, os.Interrupt)

		processes.Init()

		if err := processes.Start(); err != nil {
			log.Fatalf("Failed to boot application: %s", err)
		}

		processes.Wait()

		return nil
	},
}

func Command() *cobra.Command {
	rootCmd.Flags().StringVar(&options.Current.Endpoint, "endpoint", options.Current.Endpoint, "target server url")
	rootCmd.Flags().DurationVar(&options.Current.Period, "period", options.Current.Period, "reporting period")

	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
