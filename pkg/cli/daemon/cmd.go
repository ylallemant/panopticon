package daemon

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/ylallemant/panopticon/pkg/process"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "daemon",
	Short: "used to monitor a host",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		list, err := process.List()
		if err != nil {
			return err
		}

		// map ages
		for _, entry := range list {
			log.Printf("\t%d\t%d\t%s\n", entry.GetPID(), entry.GetPPID(), entry.GetCommand())

			// do os.* stuff on the pid
		}

		return nil
	},
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
