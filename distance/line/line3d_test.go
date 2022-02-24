package line

import (
	"Homework4/distance/point"
	"Homework4/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLine3d_Distance(t *testing.T) {

	type parameters struct {
		point1 point.Point3d
		point2 point.Point3d
	}
	type testCase struct {
		name   string
		params parameters
		want   float64
	}

	tests := []testCase{
		{
			"zero distance", parameters{*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(0, 0, 0)}, 0,
		},
		{
			"one distance", parameters{*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(1, 0, 0)}, 1,
		},
		{
			"any distance", parameters{*point.NewPoint3d(-4, 0, 2), *point.NewPoint3d(1, 3, 4)}, 6.164414002968976,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLine := mock.NewMockLine(ctrl)
	for _, tc := range tests {
		mockLine.EXPECT().Distance().Return(tc.want)
		line := NewLine3d(tc.params.point1, tc.params.point2)
		dist := line.Distance()
		assert.Equal(t, dist, mockLine.Distance())

	}
}
func TestLine3d_Coordinates(t *testing.T) {

	type parameters struct {
		point1 point.Point3d
		point2 point.Point3d
	}
	type testCase struct {
		name   string
		params parameters
		want   [][]float64
	}

	tests := []testCase{
		{
			"zero coordinates", parameters{*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(0, 0, 0)}, [][]float64{{0, 0, 0}, {0, 0, 0}},
		},
		{
			"one coordinate", parameters{*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(0, 0, 1)}, [][]float64{{0, 0, 0}, {0, 0, 1}},
		},
		{
			"any distance", parameters{*point.NewPoint3d(-10, 100, 1.3), *point.NewPoint3d(3.3, -1.2, 0)}, [][]float64{{-10, 100, 1.3}, {3.3, -1.2, 0}},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLine := mock.NewMockLine(ctrl)
	for _, tc := range tests {
		mockLine.EXPECT().Coordinates().Return(tc.want)
		line := NewLine3d(tc.params.point1, tc.params.point2)
		coordinates := line.Coordinates()
		assert.Equal(t, coordinates, mockLine.Coordinates())

	}
}
func TestLine3d_Point(t *testing.T) {
	type parameters struct {
		point1 point.Point3d
		point2 point.Point3d
	}
	type testCase struct {
		name   string
		params parameters
		want   point.Point3d
	}
	tests := []testCase{
		{
			"zero coordinates p1", parameters{*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(0, 0, 0)}, *point.NewPoint3d(0, 0, 0),
		},
		{
			"one coordinate p1", parameters{*point.NewPoint3d(0, 0, 1), *point.NewPoint3d(0, 0, 0)}, *point.NewPoint3d(0, 0, 1),
		},
		{
			"any distance p1", parameters{*point.NewPoint3d(-5.1, 0, 34), *point.NewPoint3d(-10, 4.5, -4)}, *point.NewPoint3d(-5.1, 0, 34),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			line := NewLine3d(tc.params.point1, tc.params.point2)
			assert.Equal(t, line.point1, tc.want)
		})
	}
	tests = []testCase{
		{
			"zero coordinates p2", parameters{*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(0, 0, 0)}, *point.NewPoint3d(0, 0, 0),
		},
		{
			"one coordinate p2", parameters{*point.NewPoint3d(0, 0, 1), *point.NewPoint3d(0, 0, 0)}, *point.NewPoint3d(0, 0, 0),
		},
		{
			"any distance p2", parameters{*point.NewPoint3d(-5.1, 0, 34), *point.NewPoint3d(-10, 4.5, -4)}, *point.NewPoint3d(-10, 4.5, -4),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			line := NewLine3d(tc.params.point1, tc.params.point2)
			assert.Equal(t, line.point2, tc.want)
		})
	}
}
