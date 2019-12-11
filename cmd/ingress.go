//nolint:gochecknoglobals,gochecknoinits
package cmd

import (
	"github.com/spf13/cobra"
)

var ingressCmd = &cobra.Command{
	Use:   "ingress",
	Short: "Configure ingress",
}

func init() {
	rootCmd.AddCommand(ingressCmd)
}
