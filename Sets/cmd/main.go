package main

import (
	sets "Sets/internal"
	"fmt"
)

func main() {
	mySetInt := sets.NewSetInt()
	fmt.Println(mySetInt)
	mySetInt.Add(1)
	mySetInt.Add(2)
	mySetInt.Add(3)
	mySetInt.Delete(2)
	//mySetInt.PrintElements()
	mySetInt_2 := sets.NewSetInt()
	mySetInt_2.Add(2)
	mySetInt_2.Add(3)
	union := mySetInt.Union(*mySetInt_2)
	intersection := mySetInt.Intersection(*mySetInt_2)
	difference := mySetInt.Difference(*mySetInt_2)
	mySetInt.PrintElements()
	fmt.Println("====")
	mySetInt_2.PrintElements()
	fmt.Println("union ====")
	union.PrintElements()
	fmt.Println("intersection ====")
	intersection.PrintElements()
	fmt.Println("difference ====")
	difference.PrintElements()
}
