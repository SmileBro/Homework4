package navigator

import (
	"Homework4/distance/navigator/infra"
	"Homework4/distance/point"
	"errors"
)

type Navigator struct {
	geocoder infra.Geocoding
}

type PathInfo struct {
	placeStart  infra.GeocodeData
	placeFinish infra.GeocodeData
}

func (p PathInfo) PlaceStart() infra.GeocodeData {
	return p.placeStart
}

func (p PathInfo) PlaceFinish() infra.GeocodeData {
	return p.placeFinish
}

func NewNavigator(geocoding infra.Geocoding) *Navigator {
	return &Navigator{geocoder: geocoding}
}

func (n Navigator) PathInfo(point1 point.Point, point2 point.Point) (info PathInfo, err error) {
	data1, err := n.geocoder.ReverseGeocode(point1)
	if err != nil {
		err = errors.New("error geocode first point")
		return PathInfo{}, err
	}

	data2, err := n.geocoder.ReverseGeocode(point2)
	if err != nil {
		err = errors.New("error geocode second point")
		return PathInfo{}, err
	}

	info = PathInfo{placeStart: data1, placeFinish: data2}
	return
}
