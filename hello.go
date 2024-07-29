package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")
	point1 := Point{1,1}
	point2 := Point{2,2}
	var distance float64 = getDistance(point1, point2)
	fmt.Printf("Расстояние между точками %.3f км\n", math.Round(distance*1000) / 1000)
	fmt.Printf("То же самое, но без округления: %g км\n", distance)
	var radius float64 = 1000
	enteringInRadiusFlg := checkEnteringInRadius(point1, point2, 1000)
	fmt.Printf("Точка входит в радиус %.3f? %t\n", radius, enteringInRadiusFlg)
	
	point3 := Point{1,2}
	poligonPoints := []Point{point1, point2, point3}
	poligon := Poligon{
		points: poligonPoints,
	}
	fmt.Printf("Полигон имеет %.d точек\n", len(poligon.points))
}

// Структура точки с координатами
// Координаты указываются в градусах
type Point struct {
	latitude float64
	longitude float64
}

// Функция перевода градусов в радианы
func toRadians(degress float64) float64 {
	return degress * math.Pi / 180
}

// Функция получения расстояния между двумя точками
func getDistance(point1 Point, point2 Point) float64 {
	var radiusOfEarth float64 = 6371 // В километрах
	lat1 := toRadians(point1.latitude)
	lon1 := toRadians(point1.longitude)
	lat2 := toRadians(point2.latitude)
	lon2 := toRadians(point2.longitude)

	distance := radiusOfEarth * 2 * math.Atan2(math.Sqrt(math.Pow(math.Sin((lat2 - lat1)/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin((lon2 - lon1)/2), 2)), math.Sqrt(1-math.Pow(math.Sin((lat2 - lat1)/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin((lon2 - lon1)/2), 2)))
	
	return distance
	}

// Функция проверки вхождения точки в радиус от другой точки
func checkEnteringInRadius(point1 Point, point2 Point, radius float64) bool {
	distance := getDistance(point1, point2)
	return (radius >= distance)
}

type Poligon struct {
	points []Point
}