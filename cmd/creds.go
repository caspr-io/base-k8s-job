//nolint:gochecknoglobals,gochecknoinits
package cmd

import (
	"github.com/caspr-io/caspr/internal/creds"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildCredsCmd())
}

func buildCredsCmd() *cobra.Command {
	var fetcher *creds.CredentialsFetcher = &creds.CredentialsFetcher{}

	var clusterID string

	cmd := &cobra.Command{
		Use:   "creds",
		Short: "Retrieve credentials for cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			bytes, err := fetcher.FetchCredentials(clusterID)
			if err != nil {
				log.Panic().Err(err).Msg("Could not fetch credentials")
			}

			cmd.Println(string(bytes))
			return nil
		},
	}
	cmd.Flags().StringVarP(&fetcher.Service, "service", "s", "", "The service IP address to connect to.")
	cmd.Flags().Int32VarP(&fetcher.ServicePort, "port", "p", -1, "The service port to connect to.")
	cmd.Flags().StringVarP(&clusterID, "cluster", "c", "", "The cluster id to fetch the credentials for")

	return cmd
}
