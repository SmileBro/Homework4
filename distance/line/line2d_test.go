package line

import (
	"Homework4/distance/point"
	"Homework4/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLine2d_Point(t *testing.T) {
	type parameters struct {
		point1 point.Point
		point2 point.Point
	}
	type testCase struct {
		name   string
		params parameters
		want   point.Point
	}
	tests := []testCase{
		{
			"zero coordinates p2", parameters{*point.NewPoint(0, 0), *point.NewPoint(0, 0)}, *point.NewPoint(0, 0),
		},
		{
			"one coordinate p2", parameters{*point.NewPoint(0, 0), *point.NewPoint(0, 1)}, *point.NewPoint(0, 0),
		},
		{
			"any distance p2", parameters{*point.NewPoint(3.3, -1.2), *point.NewPoint(50, 1)}, *point.NewPoint(3.3, -1.2),
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLine := mock.NewMockLine(ctrl)
	for _, tc := range tests {
		mockLine.EXPECT().Point1().Return(tc.params.point1)
		mockLine.EXPECT().Point2().Return(tc.params.point2)
		line := NewLine2d(tc.params.point1, tc.params.point2)
		assert.Equal(t, line.Point1(), mockLine.Point1())
		assert.Equal(t, line.Point2(), mockLine.Point2())
	}
}
func TestLine2d_Coordinates(t *testing.T) {

	type parameters struct {
		point1 point.Point
		point2 point.Point
	}
	type testCase struct {
		name   string
		params parameters
		want   [][]float64
	}

	tests := []testCase{
		{
			"zero coordinates", parameters{*point.NewPoint(0, 0), *point.NewPoint(0, 0)}, [][]float64{{0, 0}, {0, 0}},
		},
		{
			"one coordinate", parameters{*point.NewPoint(0, 0), *point.NewPoint(0, 1)}, [][]float64{{0, 0}, {0, 1}},
		},
		{
			"any distance", parameters{*point.NewPoint(-10, 100), *point.NewPoint(3.3, -1.2)}, [][]float64{{-10, 100}, {3.3, -1.2}},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	for _, tc := range tests {
		mockLine := mock.NewMockLine(ctrl)
		mockLine.EXPECT().Coordinates().Return(tc.want)
		line := NewLine2d(tc.params.point1, tc.params.point2)
		coordinates := line.Coordinates()
		assert.Equal(t, coordinates, mockLine.Coordinates())

	}
}
func TestLine2d_Distance(t *testing.T) {

	type parameters struct {
		point1 point.Point
		point2 point.Point
	}
	type testCase struct {
		name   string
		params parameters
		want   float64
	}

	tests := []testCase{
		{
			"zero distance", parameters{*point.NewPoint(0, 0), *point.NewPoint(0, 0)}, 0,
		},
		{
			"one distance", parameters{*point.NewPoint(0, 0), *point.NewPoint(0, 1)}, 1,
		},
		{
			"any distance", parameters{*point.NewPoint(-1, -1), *point.NewPoint(1, 1)}, 2.8284271247461903,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLine := mock.NewMockLine(ctrl)
	for _, tc := range tests {
		mockLine.EXPECT().Distance().Return(tc.want)
		line := NewLine2d(tc.params.point1, tc.params.point2)
		assert.Equal(t, line.Distance(), mockLine.Distance())
	}
}
