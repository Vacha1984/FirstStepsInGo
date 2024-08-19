package points

import (
	"errors"
	"fmt"
)

type Poligon struct {
	points []Point
}

func NewPoligon(p []Point) (Poligon, error) {
	if len(p) < 3 {
		return Poligon{}, errors.New("Invalid poligon")
	}
	return Poligon{points: p}, nil
}

func (p Poligon) PrintPoints() {
	for _, value := range p.points {
		fmt.Printf("(%g, %g)\n", value.latitude, value.longitude)
	}
}
