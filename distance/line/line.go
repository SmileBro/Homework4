package line

import "Homework4/distance/point"

type Line interface {
	Distance() float64
	Coordinates() [][]float64
	Point1() point.Point
	Point2() point.Point
	//MarshalJSON() ([]byte, error)
}
