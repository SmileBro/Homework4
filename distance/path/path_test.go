package path

import (
	"Homework4/distance/line"
	"Homework4/distance/point"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPath_Distance(t *testing.T) {
	type parameters struct {
		lines []line.Line
	}
	type testCase struct {
		name   string
		params parameters
		want   float64
	}
	tests := []testCase{
		{
			"one line", parameters{[]line.Line{line.NewLine2d(*point.NewPoint(0, 0), *point.NewPoint(0, 0))}}, 0,
		},
		{
			"two lines", parameters{[]line.Line{line.NewLine2d(*point.NewPoint(0, 0), *point.NewPoint(0, 0)), line.NewLine2d(*point.NewPoint(0, 0), *point.NewPoint(0, 1))}}, 1,
		},
		{
			"two lines any distance", parameters{[]line.Line{line.NewLine2d(*point.NewPoint(-2.3, 4), *point.NewPoint(1.2, -4)), line.NewLine2d(*point.NewPoint(1.2, -4), *point.NewPoint(-8, 10))}}, 25.484437871384028,
		},
		{
			"one line3d", parameters{[]line.Line{line.NewLine3d(*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(0, 0, 0))}}, 0,
		},
		{
			"two line3d", parameters{[]line.Line{line.NewLine3d(*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(0, 0, 0)), line.NewLine3d(*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(0, 1, 0))}}, 1,
		},
		{
			"no lines", parameters{[]line.Line{}}, 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lines := NewPath(tc.params.lines)
			dist, err := lines.Distance()
			if err != nil && dist != 0 {
				assert.Nil(t, err)
			}
			assert.Equal(t, dist, tc.want)
		})
	}
}
