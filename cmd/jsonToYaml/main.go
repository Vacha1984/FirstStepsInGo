package main

import (
	"example/hello/internal/jsonToYaml"
	"fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
)

func main() {
	fileFlag := pflag.StringP("file", "f", "", "file to parse")
	pflag.Parse()
	pathToFile := *fileFlag
	fmt.Println("Получено значение:", pathToFile)

	fileContent, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		fmt.Println(err)
	}

	fileData := string(fileContent)
	fmt.Println(fileData)
	jsonToYaml.PareserJson(fileData)
}
