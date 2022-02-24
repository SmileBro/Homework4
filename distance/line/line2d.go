package line

import (
	"Homework4/distance/point"
	"math"
)

type Line2d struct {
	point1 point.Point
	point2 point.Point
}

/*func (l Line2d) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(
		struct {
			Point1 point.Point `json:"p1"`
			Point2 point.Point `json:"p2"`
		}{l.point1, l.point2})

	if err != nil {
		return nil, err
	}

	return j, err
}*/

func (l Line2d) Point1() point.Point {
	return l.point1
}

func (l Line2d) Point2() point.Point {
	return l.point2
}

func NewLine2d(point1 point.Point, point2 point.Point) *Line2d {
	return &Line2d{point1: point1, point2: point2}
}

func (l Line2d) Distance() (distance float64) {
	distance = math.Sqrt(math.Pow(l.point2.X()-l.point1.X(), 2) + math.Pow(l.point2.Y()-l.point1.Y(), 2))
	return
}
func (l Line2d) Coordinates() [][]float64 {
	coordinates := make([][]float64, 2)
	coordinates[0] = append(coordinates[0], l.point1.X())
	coordinates[0] = append(coordinates[0], l.point1.Y())
	coordinates[1] = append(coordinates[1], l.point2.X())
	coordinates[1] = append(coordinates[1], l.point2.Y())

	return coordinates
}
