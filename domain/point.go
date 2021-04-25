package domain

import "math"

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p Point) DistanceTo(p2 Point) int {
	return int(math.Abs(float64(p2.X-p.X)) + math.Abs(float64(p2.Y-p.Y)))
}
