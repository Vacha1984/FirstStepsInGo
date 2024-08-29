package jsonToYaml

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPareserJson(t *testing.T) {
	json_text := "{\"tags\": {\"tagName\" : \"test\"}}"
	json_byte := []byte(json_text)
	yaml_result := PareserJson(json_byte)
	mapForYaml := make(map[string]map[string]interface{})
	yaml.Unmarshal(yaml_result, &mapForYaml)
	fmt.Print(mapForYaml)
	assert.Equal(t, mapForYaml["tags"]["default"], "", "%v != \"\"", mapForYaml["tags"]["default"])

	type args struct {
		jsonStr []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PareserJson(tt.args.jsonStr)
		})
	}
}
