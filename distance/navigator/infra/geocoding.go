package infra

import "Homework4/distance/point"

type Geocoding interface {
	Geocode(address string) (point point.Point, err error)
	ReverseGeocode(point point.Point) (data GeocodeData, err error)
}

type GeocodeData struct {
	Point      point.Point
	Address    string
	PostalCode string
}
