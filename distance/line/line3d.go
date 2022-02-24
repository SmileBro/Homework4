package line

import (
	"Homework4/distance/point"
	"math"
)

type Line3d struct {
	point1 point.Point3d
	point2 point.Point3d
}

/*func (l Line3d) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(
		struct {
			Point1 point.Point3d `json:"p1"`
			Point2 point.Point3d `json:"p2"`
		}{l.point1, l.point2})

	if err != nil {
		return nil, err
	}

	return j, err
}*/

func (l Line3d) Point1() point.Point3d {
	return l.point1
}

func (l Line3d) Point2() point.Point3d {
	return l.point2
}

func NewLine3d(point1 point.Point3d, point2 point.Point3d) *Line3d {
	return &Line3d{point1: point1, point2: point2}
}

// = √(xb - xa)2 + (yb - ya)2 + (zb - za)2
func (l Line3d) Distance() (distance float64) {
	distance = math.Sqrt(math.Pow(l.point2.X()-l.point1.X(), 2) + math.Pow(l.point2.Y()-l.point1.Y(), 2) + math.Pow(l.point2.Z()-l.point1.Z(), 2))
	return
}
func (l Line3d) Coordinates() [][]float64 {
	coordinates := make([][]float64, 2)
	coordinates[0] = append(coordinates[0], l.point1.X())
	coordinates[0] = append(coordinates[0], l.point1.Y())
	coordinates[0] = append(coordinates[0], l.point1.Z())
	coordinates[1] = append(coordinates[1], l.point2.X())
	coordinates[1] = append(coordinates[1], l.point2.Y())
	coordinates[1] = append(coordinates[1], l.point2.Z())
	return coordinates
}
