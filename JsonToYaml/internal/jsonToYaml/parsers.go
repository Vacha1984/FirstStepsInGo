package jsonToYaml

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"reflect"
)

func PareserJson(jsonStr []byte) []byte {
	var jsonAsStruct interface{}
	err := json.Unmarshal(jsonStr, &jsonAsStruct)
	if err != nil {
	}
	fmt.Println(jsonAsStruct)
	newMapForYaml := make(map[string]interface{})
	for key, value := range jsonAsStruct.(map[string]interface{}) {
		tmpMapsForYaml := []map[string]interface{}{}
		for key, value := range value.(map[string]interface{}) {
			tmpMapForYaml := make(map[string]interface{})
			tmpMapForYaml["name"] = key
			tmpMapForYaml["type"] = reflect.TypeOf(value).Name()
			tmpMapForYaml["default"] = getZeRoValue(value)
			tmpMapsForYaml = append(tmpMapsForYaml, tmpMapForYaml)
		}
		newMapForYaml[key] = tmpMapsForYaml
	}
	fmt.Println(newMapForYaml)
	newMapAsYaml, err := yaml.Marshal(&newMapForYaml)
	if err != nil {
	}
	fmt.Println(string(newMapAsYaml))
	return newMapAsYaml
}

func getZeRoValue(value interface{}) interface{} {
	switch value.(type) {
	case int:
		var x int
		return x
	case float64:
		var x float64
		return x
	case bool:
		var x bool
		return x
	case string:
		var x string
		return x
	default:
		return "Тип не распознан"
	}
}
