package sets

import "fmt"

type SetInt struct {
	Elements map[int]struct{}
}

func NewSetInt() *SetInt {
	var setInt SetInt
	setInt.Elements = make(map[int]struct{})
	return &setInt
}

func (set *SetInt) Add(element int) {
	set.Elements[element] = struct{}{}
}

func (set *SetInt) Delete(element int) {
	delete(set.Elements, element)
}

func (set *SetInt) PrintElements() {
	for element := range set.Elements {
		fmt.Println(element)
	}
}

func (set SetInt) Union(otherSet SetInt) *SetInt {
	newSet := NewSetInt()
	for element := range otherSet.Elements {
		newSet.Elements[element] = struct{}{}
	}
	for element := range set.Elements {
		newSet.Elements[element] = struct{}{}
	}
	return newSet
}

func (set SetInt) Intersection(otherSet SetInt) *SetInt {
	newSet := NewSetInt()
	for element := range set.Elements {
		_, exists := otherSet.Elements[element]
		if exists {
			newSet.Elements[element] = struct{}{}
		}
	}
	return newSet
}

func (set SetInt) Difference(otherSet SetInt) *SetInt {
	newSet := NewSetInt()
	for element := range set.Elements {
		if _, exists := otherSet.Elements[element]; !exists {
			newSet.Elements[element] = struct{}{}
		}
	}
	return newSet
}
