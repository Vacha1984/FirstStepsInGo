package points

import (
	"errors"
	"math"
)

type degrees float64

type Point struct {
	latitude, longitude degrees
}

func NewPoint(lat, lon degrees) (Point, error) {
	if lat <= 90 && lat >= -90 && lon <= 180 && lon >= -180 {
		return Point{latitude: lat, longitude: lon}, nil
	}
	return Point{}, errors.New("invalid point")
}

func (p1 Point) DistanceTo(p2 Point) uint32 {
	var radiusOfEarth float64 = 6371000 // В метрах
	lat1 := p1.latitude.asRadians()
	lat2 := p2.latitude.asRadians()
	lon1 := p1.longitude.asRadians()
	lon2 := p2.longitude.asRadians()
	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := radiusOfEarth * c
	return uint32(distance)
}

func (deg degrees) asRadians() float64 {
	return float64(deg) * math.Pi / 180
}

func (p1 Point) CheckEnteringInRadius(p2 Point, radius uint32) bool {
	return p1.DistanceTo(p2) > radius
}
