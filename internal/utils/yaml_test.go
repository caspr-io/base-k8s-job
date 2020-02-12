package utils

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestYamlPathCorrect(t *testing.T) {
	yaml := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": "baz",
			"boo": "beh",
		},
	}

	v, err := YamlPath(yaml, "foo.bar")
	assert.NilError(t, err)
	assert.Equal(t, v, "baz")
}
