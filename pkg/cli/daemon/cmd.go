package daemon

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	daemonOptions "github.com/ylallemant/panopticon/pkg/cli/daemon/options"
	serverOptions "github.com/ylallemant/panopticon/pkg/cli/server/options"
	"github.com/ylallemant/panopticon/pkg/daemon"
	"github.com/ylallemant/panopticon/pkg/server"
	"github.com/ylallemant/panopticon/pkg/server/service/graceful"
)

var rootCmd = &cobra.Command{
	Use:   "daemon",
	Short: "used to monitor a host",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		startDefaultServer := false

		if daemonOptions.Current.ConfigPath != "" {
			fileOptions, err := daemonOptions.Current.Load(daemonOptions.Current.ConfigPath)
			if err != nil {
				return err
			}

			daemonOptions.Current = fileOptions
		}

		if daemonOptions.Current.Endpoint == "" {
			daemonOptions.Current.Endpoint = fmt.Sprintf("grpc://localhost:%d", serverOptions.Current.Ports.GRPC)
			startDefaultServer = true
		}

		processes := graceful.NewProcessBucket()

		deamon, err := daemon.NewDaemon(daemonOptions.Current)
		if err != nil {
			return err
		}

		processes.AddProcess(deamon)

		if startDefaultServer {
			serverOptions := serverOptions.Current
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
	rootCmd.Flags().StringVar(&daemonOptions.Current.ConfigPath, "config-path", daemonOptions.Current.ConfigPath, "path to configuration file")
	rootCmd.Flags().StringVar(&daemonOptions.Current.Endpoint, "endpoint", daemonOptions.Current.Endpoint, "target server url")
	rootCmd.Flags().DurationVar(&daemonOptions.Current.Period, "period", daemonOptions.Current.Period, "reporting period")

	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
