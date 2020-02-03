package utils

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"

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

func ReadYaml(file string) (map[string]interface{}, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := yaml.NewDecoder(reader)

	var yamlContents map[string]interface{}

	err = decoder.Decode(yamlContents)
	if err != nil {
		return nil, err
	}

	return yamlContents, nil
}

func WriteYaml(file string, contents map[string]interface{}) error {
	c := ToYaml(contents)

	return ioutil.WriteFile(file, c, 0644)
}
