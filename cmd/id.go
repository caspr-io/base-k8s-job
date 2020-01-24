//nolint:gochecknoglobals,gochecknoinits
package cmd

import (
	"os"

	"github.com/caspr-io/mu-kit/id"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildIDCmd())
}

func buildIDCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "id",
		Short: "Returns a new shortId with the given prefix",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(id.New(args[0]))
		},
	}

	cmd.SetOut(os.Stdout)

	return cmd
}
