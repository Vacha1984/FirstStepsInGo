package jsonToYaml

import (
	"fmt"
	"strconv"
	"strings"
)

func PareserJson(jsonStr string) {
	splitedJson := strings.Split(jsonStr, "{")
	// Сильно упрощено, используя последний элемент сплита,
	// тк принимаю за истину что в файле нет сложной вложенности
	tagAndValue := strings.Split(splitedJson[len(splitedJson)-1], ",")
	mapTagValues := make(map[string]string)
	for _, item := range tagAndValue {
		splitedItem := strings.Split(item, ":")
		mapTagValues[splitedItem[0]] = splitedItem[1]
	}

	mapTagType := make(map[string]interface{})
	for key, valueWithSpace := range mapTagValues {
		value := strings.ReplaceAll(valueWithSpace, " ", "")
		if strings.Contains(value, "\"") {
			mapTagType[key] = ""
			continue
		}
		_, err := strconv.ParseBool(value)
		if err == nil {
			mapTagType[key] = false
			continue
		}
		_, err = strconv.ParseInt(value, 10, 64)
		if err == nil {
			mapTagType[key] = 0
			continue
		}
		_, err = strconv.ParseFloat(value, 64)
		if err == nil {
			mapTagType[key] = 0.0
			continue
		}
	}

	for key, value := range mapTagType {
		fmt.Printf("%s имеет значение %v типа %T", key, value, value)
	}
	fmt.Print("\n")
}
