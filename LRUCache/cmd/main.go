package main

import (
	LRUCache "LRUCache/internal"
	"fmt"
)

func main() {
	cache := LRUCache.NewLRUCacheAdd(3)
	x := cache.AddWithCache([2]int{1, 2})
	y := cache.AddWithCache([2]int{1, 3})
	z := cache.AddWithCache([2]int{2, 3})
	xx := cache.AddWithCache([2]int{1, 2})
	zz := cache.AddWithCache([2]int{2, 3})
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(xx)
	fmt.Println(zz)

}
