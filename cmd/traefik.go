//nolint:gochecknoglobals,gochecknoinits
package cmd

import (
	"github.com/caspr-io/caspr/internal/traefik"
	"github.com/spf13/cobra"
)

var ingress *traefik.Ingress = &traefik.Ingress{}

func init() {
	traefikCmd.Flags().StringVarP(&ingress.Namespace, "namespace", "n", "default", "The namespace to add the ingress to")
	traefikCmd.Flags().StringVarP(&ingress.Service, "service", "s", "", "The service to add the ingress for")
	traefikCmd.Flags().Int32VarP(&ingress.Port, "port", "p", -1, "The port of the service to add the ingress for")
	traefikCmd.Flags().BoolVarP(&ingress.TLS, "tls", "t", false, "Whether the service should be exposed on TLS")
	traefikCmd.Flags().StringToStringVarP(&ingress.Labels, "label", "l", map[string]string{}, "Labels to define on the K8s object(s)")
	traefikCmd.Flags().StringVarP(&ingress.URL, "url", "u", "", "The (full) URL on which the root of the application will be served")
	traefikCmd.MarkFlagRequired("service") //nolint:errcheck
	traefikCmd.MarkFlagRequired("port")    //nolint:errcheck
	traefikCmd.MarkFlagRequired("url")     //nolint:errcheck
	ingressCmd.AddCommand(traefikCmd)
}

var traefikCmd = &cobra.Command{
	Use:   "traefik",
	Short: "Configure ingress using Traefik",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := traefik.NewTraefik()
		if err != nil {
			return err
		}

		if err := t.CreateIngress(ingress); err != nil {
			return err
		}

		return nil
	},
}
