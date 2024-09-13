package packUnpackStrings

import (
	"fmt"
	"strconv"
	"strings"
)

func Unpack(s string) string {
	numbers := "1234567890"
	currentNumber := make([]string, 0, 5)
	newString := make([]string, 5, 10)
	var currentChar string
	for index, char := range s {
		// Eсли первый элемент сразу число
		if index == 0 && strings.Contains(numbers, string(char)) {
			fmt.Println("Случай 1")
			return ""
		}
		// Если это число
		if strings.Contains(numbers, string(char)) {
			fmt.Println("Случай 2")
			currentNumber = append(currentNumber, string(char))
			continue
		}
		// Если обычный символ и до этого был символ
		if !strings.Contains(numbers, string(char)) && len(currentNumber) == 0 {
			fmt.Println("Случай 3")
			newString = append(newString, currentChar)
			currentChar = string(char)
			continue
		}
		// Если число закончилось
		if !strings.Contains(numbers, string(char)) && len(currentNumber) != 0 {
			fmt.Println("Случай 4")
			intNumber, err := strArrayToInt(currentNumber)
			if err == nil {
				newString = append(newString, strings.Repeat(currentChar, intNumber))
				currentNumber = nil
				currentChar = string(char)
			}
		}
	}
	if len(currentNumber) != 0 {
		fmt.Println("Случай 5")
		intNumber, err := strArrayToInt(currentNumber)
		if err == nil {
			newString = append(newString, strings.Repeat(currentChar, intNumber))
		}
	} else {
		newString = append(newString, currentChar)
	}
	return strings.Join(newString, "")
}

func strArrayToInt(array []string) (int, error) {
	strNumber := strings.Join(array, "")
	return strconv.Atoi(strNumber)

}
