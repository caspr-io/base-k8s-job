//nolint:gochecknoinits
package cmd

import (
	"github.com/caspr-io/caspr/internal/kubeconfig"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildKubeconfigCmd())
}

func buildKubeconfigCmd() *cobra.Command {
	var file string

	var outputFile string

	var configType string

	var targetNamespace string

	cmd := &cobra.Command{
		Use:   "kubeconfig",
		Short: "Build a kubeconfig file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return kubeconfig.NewBuilder(configType, file, targetNamespace).Build(outputFile)
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "/caspr-credentials/creds.yaml", "The file to read the K8s credentials from")
	cmd.Flags().StringVarP(&outputFile, "output", "o", "k8s.yaml", "The output file for the kubernetes config")
	cmd.Flags().StringVarP(&configType, "type", "t", "", "The type of cluster that the credentials are for.")
	cmd.Flags().StringVarP(&targetNamespace, "namespace", "n", "", "The namespace in the cluster to connect to.")

	return cmd
}
