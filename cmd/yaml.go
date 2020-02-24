//nolint:gochecknoinits
package cmd

import (
	"github.com/caspr-io/caspr/internal/utils"
	"github.com/caspr-io/yamlpath"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func buildYamlCmd() *cobra.Command {
	var file string

	cmd := &cobra.Command{
		Use:   "yaml",
		Short: "Get value from YAML",
		Args:  cobra.ExactArgs(1), //nolint:gomnd
		RunE: func(cmd *cobra.Command, args []string) error {
			contents, err := utils.ReadYaml(file)
			if err != nil {
				return err
			}

			result, err := yamlpath.YamlPath(contents, args[0])
			if err != nil {
				return err
			}

			bytes, err := yaml.Marshal(result)
			if err != nil {
				return err
			}

			println(string(bytes))

			return nil
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "The YAML file to parse")
	cmd.MarkFlagRequired("file") //nolint:errcheck

	return cmd
}

func init() {
	rootCmd.AddCommand(buildYamlCmd())
}
