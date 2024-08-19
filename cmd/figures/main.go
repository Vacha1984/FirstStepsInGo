package main

import (
	"example/hello/internal/points"
	"fmt"
)

func main() {
	fmt.Println(`Hello, World!`)
	point1, _ := points.NewPoint(1, 1)
	point2, _ := points.NewPoint(2, 2)
	distance := point1.DistanceTo(point2)
	fmt.Printf("Расстояние между точками %d м.\n", distance)

	var radius uint32 = 1000
	enteringInRadiusFlg := point1.CheckEnteringInRadius(point2, radius)
	fmt.Printf("Точка входит в радиус %d м? %t\n", radius, enteringInRadiusFlg)

	point3, _ := points.NewPoint(1, 1)
	poligon1, _ := points.NewPoligon([]points.Point{point1, point2, point3})
	poligon1.PrintPoints()
}
