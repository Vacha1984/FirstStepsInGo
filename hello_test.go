package main

import "testing"

func TestGetDistancePositivePosirive(t *testing.T) {
	point1 := Point{1,1}
	point2 := Point{2,2}
	result := getDistance(point1, point2)
	var expected float64 = 157.2134688151181
	
	if result != expected {
		t.Errorf("%f != %f", result, expected)
	}
}
	
func TestGetDistanceNegativeNegative(t *testing.T) {
	point1 := Point{-1,-1}
	point2 := Point{-2,-2}
	result := getDistance(point1, point2)
	var expected float64 = 157.2134688151181
	
	if result != expected {
		t.Errorf("%f != %f", result, expected)
	}
}

func TestGetDistanceBoundary(t *testing.T) {
	point1 := Point{90,180}
	point2 := Point{-90,-180}
	result := getDistance(point1, point2)
	var expected float64 = 20015.086796020572
	
	if result != expected {
		t.Errorf("%f != %f", result, expected)
	}
}

func TestCheckEnteringInRadiusPositive(t *testing.T) {
	point1 := Point{1,1}
	point2 := Point{2,2}
	var radius float64 = 300
	result := checkEnteringInRadius(point1, point2, radius)
	var expected bool = true
	
	if result != expected {
		t.Errorf("%t != %t", result, expected)
	}
}

func TestCheckEnteringInRadiusNegative(t *testing.T) {
	point1 := Point{1,1}
	point2 := Point{2,2}
	var radius float64 = 100
	result := checkEnteringInRadius(point1, point2, radius)
	var expected bool = false
	
	if result != expected {
		t.Errorf("%t != %t", result, expected)
	}
}

func TestCheckEnteringInRadiusBoundary(t *testing.T) {
	point1 := Point{1,1}
	point2 := Point{2,2}
	var radius float64 = 157.2134688151181
	result := checkEnteringInRadius(point1, point2, radius)
	var expected bool = true
	
	if result != expected {
		t.Errorf("%t != %t", result, expected)
	}
}