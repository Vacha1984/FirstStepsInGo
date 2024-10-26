package LRUCache

import (
	"container/list"
	"fmt"
)

type LRUCacheAddItem struct {
	key   [2]int
	value int
}

type LRUCacheAdd struct {
	capacity int
	cache    map[[2]int]*list.Element
	list     *list.List
}

func NewLRUCacheAdd(capacity int) *LRUCacheAdd {
	return &LRUCacheAdd{
		capacity: capacity,
		cache:    make(map[[2]int]*list.Element),
		list:     list.New(),
	}
}

func Add(x, y int) int {
	return x + y
}

func (c *LRUCacheAdd) AddWithCache(key [2]int) any {
	if element, exists := c.cache[key]; exists {
		value := element.Value.(LRUCacheAddItem).value
		c.list.MoveToFront(element)
		fmt.Println("Значение взято из кэша: ", value)
		return value
	}
	if c.list.Len() >= c.capacity {
		back := c.list.Back()
		if back != nil {
			c.list.Remove(back)
			delete(c.cache, back.Value.(LRUCacheAddItem).key)
		}
	}
	value := Add(key[0], key[1])
	newItem := LRUCacheAddItem{key, value}
	element := c.list.PushFront(newItem)
	c.cache[key] = element
	fmt.Println("Значение вычеслено заново: ", value)
	return value
}
