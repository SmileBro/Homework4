package path

import (
	"Homework4/distance/line"
	"errors"
	"fmt"
)

/*type Path interface {
	Distance() float64
}
*/
type Path struct {
	lines []line.Line
}

func NewPath(lines []line.Line) *Path {
	return &Path{lines: lines}
}

func (l *Path) AddLine(d line.Line) *Path {
	l.lines = append(l.lines, d)
	return l
}

func (l Path) Distance() (distance float64, err error) {

	if len(l.lines) < 1 {
		return distance, errors.New("путь пустой")
	}

	for _, p := range l.lines {
		distanceTwoPoint := p.Distance()
		fmt.Println("Растояние между точками", p.Coordinates(), " равно ", distanceTwoPoint)
		distance += distanceTwoPoint
	}
	return
}
