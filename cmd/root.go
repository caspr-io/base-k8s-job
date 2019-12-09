package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/caspr-io/caspr-result/internal"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var file string

var rootCmd = &cobra.Command{
	Use:   "caspr-result",
	Short: "Sends the provisioning result back to CASPR",
	Run: func(cmd *cobra.Command, args []string) {
		if strings.TrimSpace(file) != "" {
			log.Logger.Info().Str("file", file).Msg("Reading payload from file")
			fileHandle, err := os.Open(file)
			if err != nil {
				panic(err)
			}
			defer fileHandle.Close()
			payloadReader := bufio.NewReader(fileHandle)
			internal.ReadPayload(payloadReader).Send()
		} else if stdInAvailable() {
			log.Logger.Info().Str("file", "<stdin>").Msg("Reading payload from os.Stdin")
			reader := bufio.NewReader(os.Stdin)
			internal.ReadPayload(reader).Send()
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "The file that contains the YAML result")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func stdInAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return stat.Mode()&os.ModeCharDevice == 0
}
