package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

	yamlContents := map[string]interface{}{}

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

func YamlPath(yaml interface{}, path string) (interface{}, error) {
	components := strings.SplitN(path, ".", 2)

	switch y := yaml.(type) {
	case map[string]interface{}:
		if v, ok := y[components[0]]; ok {
			if len(components) > 1 {
				return YamlPath(v, components[1])
			}

			return v, nil
		}

		return nil, fmt.Errorf("could not find key %s in YAML", components[0])
	default:
		return nil, fmt.Errorf("could not walk path '%s' as it is not a map %v", components[0], y)
	}
}
