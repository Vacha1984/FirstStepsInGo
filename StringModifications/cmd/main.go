package main

import (
	"example/stringModifications/internal/packUnpackStrings"
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	strFlag := pflag.StringP("string", "s", "", "string to modification")
	pflag.Parse()
	strToModification := *strFlag
	unpacked_str := packUnpackStrings.Unpack(strToModification)
	fmt.Println(unpacked_str)
}
