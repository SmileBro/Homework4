package navigator

import (
	"Homework4/distance/navigator/infra"
	"Homework4/distance/point"
	"Homework4/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNavigator_PathInfo(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGeocoder := mock.NewMockGeocoding(ctrl)
	point1 := point.NewPoint(0, 0)
	point2 := point.NewPoint(1, 1)
	gomock.InOrder(
		mockGeocoder.EXPECT().ReverseGeocode(*point1).Return(infra.GeocodeData{Point: *point1, Address: "first point address", PostalCode: "123321"}, nil),
		mockGeocoder.EXPECT().ReverseGeocode(*point2).Return(infra.GeocodeData{Point: *point2, Address: "second point address", PostalCode: "111321"}, nil),
	)

	navi := NewNavigator(mockGeocoder)
	info, err := navi.PathInfo(*point1, *point2)
	assert.Nil(t, err)
	assert.Equal(t, info.PlaceStart().Point, *point1)
	assert.Equal(t, info.PlaceFinish().Point, *point2)
	assert.Equal(t, info.PlaceFinish().PostalCode, "111321")
	assert.Equal(t, info.PlaceStart().PostalCode, "123321")

}
