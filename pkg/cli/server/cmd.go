package server

import (
	"github.com/spf13/pflag"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "used to centralise host monitoring",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
