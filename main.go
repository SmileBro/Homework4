package main

import (
	"Homework4/distance/point"
	geocoder2 "Homework4/geocoder"
	"fmt"
)

func main() {
	line := "москва"
	p := *point.NewPoint(55.601983, 37.359486)
	geocoder := geocoder2.NewGeocoder()
	fmt.Println(geocoder.Geocode(line))
	fmt.Println(geocoder.ReverseGeocode(p))

}

/*func PrintfLn(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format+"\n", a...)
}*/
