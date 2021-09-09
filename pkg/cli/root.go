package cli

import (
	"github.com/spf13/cobra"
	"github.com/ylallemant/panopticon/pkg/cli/daemon"
	"github.com/ylallemant/panopticon/pkg/cli/server"
	"github.com/ylallemant/panopticon/pkg/process"
)

var rootCmd = &cobra.Command{
	Use:   "panopticon",
	Short: "panopticon contains a collection of cli tools",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		_, err := process.List()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(daemon.Command())
	rootCmd.AddCommand(server.Command())
}

func Command() *cobra.Command {
	return rootCmd
}
