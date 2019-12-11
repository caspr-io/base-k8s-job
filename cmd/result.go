//nolint:gochecknoglobals,gochecknoinits
package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/caspr-io/caspr/internal/result"
	"github.com/caspr-io/caspr/internal/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var file string

var resultCmd = &cobra.Command{
	Use:   "result",
	Short: "Send the provisioning result back to CASPR",
	Run: func(cmd *cobra.Command, args []string) {
		if strings.TrimSpace(file) != "" {
			log.Logger.Info().Str("file", file).Msg("Reading payload from file")
			fileHandle, err := os.Open(file)
			if err != nil {
				panic(err)
			}
			defer fileHandle.Close()
			payloadReader := bufio.NewReader(fileHandle)
			result.ReadPayload(payloadReader).Send()
		} else if utils.StdinAvailable() {
			log.Logger.Info().Str("file", "<stdin>").Msg("Reading payload from os.Stdin")
			reader := bufio.NewReader(os.Stdin)
			result.ReadPayload(reader).Send()
		}
	},
}

func init() {
	rootCmd.AddCommand(resultCmd)
	resultCmd.Flags().StringVarP(&file, "file", "f", "", "The file that contains the YAML result")
}
