package geocoder

import (
	"Homework4/distance/navigator/infra"
	point2 "Homework4/distance/point"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Geocoder struct {
	client *http.Client
	urlr   string
	urlg   string
	token  string
}
type GeoInfo struct {
	Source       string `json:"source"`
	Result       string `json:"result"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Region       string `json:"region"`
	CityArea     string `json:"city_area"`
	CityDistrict string `json:"city_district"`
	Street       string `json:"street"`
	House        string `json:"house"`
	GeoLat       string `json:"geo_lat"`
	GeoLon       string `json:"geo_lon"`
	QcGeo        int    `json:"qc_geo"`
}

func (g Geocoder) Geocode(address string) (point point2.Point, err error) {
	jsonRequest, _ := json.Marshal([]string{address})
	req, _ := http.NewRequest("POST", g.urlg, bytes.NewBuffer(jsonRequest))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Token "+g.token)
	req.Header.Add("X-Secret", "3879e1f0acb8363a85311c008de0c810dc8474a4")
	response, err := g.client.Do(req)
	if err != nil {
		return
	}
	var res GeoInfo
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v/n", res)
	return
}

func (g Geocoder) ReverseGeocode(point point2.Point) (data infra.GeocodeData, err error) {
	jsonRequest, _ := json.Marshal(map[string]string{"lat": fmt.Sprintf("%f", point.X()), "lon": fmt.Sprintf("%f", point.Y()), "count": "1"})
	req, _ := http.NewRequest("POST", g.urlr, bytes.NewBuffer(jsonRequest))
	req.Header.Add("Authorization", "Token "+g.token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	response, err := g.client.Do(req)
	if err != nil {
		return
	}

	var res map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return
	}
	var suggestions = res["suggestions"].([]interface{})
	sugElem := suggestions[0].(map[string]interface{})
	data1 := sugElem["data"].(map[string]interface{})
	//postalCode := data1["postal_code"].(string)
	fmt.Println(data1)

	return
}

/*
curl -X POST \
-H "Content-Type: application/json" \
-H "Accept: application/json" \
-H "Authorization: Token 11704dfd376beff80fa6b77d454026c9ddd025f0" \
-d '{ "lat": 55.878, "lon": 37.653 }' \
https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address
*/

func NewGeocoder() *Geocoder {
	//return &Geocoder{client: &http.Client{}, url: "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", token: "1409412d7fd853b91c7fc572d5775639e958da26"}
	return &Geocoder{client: &http.Client{}, urlr: "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", urlg: "https://cleaner.dadata.ru/api/v1/clean/address", token: "1409412d7fd853b91c7fc572d5775639e958da26"}
}
