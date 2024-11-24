package main

import (
	pi "PI/internal"
	"fmt"
)

func main() {
	//gorutinsAmount := pflag.IntP("gorutins_amount", "n", 1, "amount of gorutins")
	c := make(chan float64, 40)
	for totalIndex := 0; totalIndex < 10; totalIndex++ {
		for gorunitIndex := 1; gorunitIndex <= 4; gorunitIndex++ {
			leibnicIndex := totalIndex*4 + gorunitIndex
			go pi.LeibnicSum(leibnicIndex, c)
		}
		sum := 0.0
		for i := range c {
			sum += i
		}
		close(c)
		fmt.Println(sum * 4)
	}
}
