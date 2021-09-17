package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/panopticon/pkg/cli/server/options"
	"github.com/ylallemant/panopticon/pkg/server"
	"github.com/ylallemant/panopticon/pkg/server/service/graceful"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "used to centralise host monitoring",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		service := server.NewServer(options.Current)

		processes := graceful.NewProcessBucket()
		processes.AddProcess(service)

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
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
