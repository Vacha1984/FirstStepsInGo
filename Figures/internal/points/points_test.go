package points

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDistancePositivePosirive(t *testing.T) {
	point1 := Point{1, 1}
	point2 := Point{2, 2}
	result := point1.DistanceTo(point2)
	assert.Equal(t, result, uint32(157225), "%d != %d", result, uint32(157225))
}

func TestGetDistanceNegativeNegative(t *testing.T) {
	point1 := Point{-1, -1}
	point2 := Point{-2, -2}
	result := point1.DistanceTo(point2)
	assert.Equal(t, result, uint32(157225), "%d != %d", result, uint32(157225))
}
