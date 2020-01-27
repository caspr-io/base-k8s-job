package utils

import (
	"bytes"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// ToYaml encodes an object to a byte array containing YAML
func ToYaml(value interface{}) []byte {
	byteBuffer := &bytes.Buffer{}

	encoder := yaml.NewEncoder(byteBuffer)
	defer encoder.Close()

	if err := encoder.Encode(value); err != nil {
		log.Panic().Err(err).Msg("Could not encode YAML")
	}

	return byteBuffer.Bytes()
}
