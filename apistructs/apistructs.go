package apistructs

import (
	"encoding/json"
	"log"
	"strings"
)

// APIKey JSON for Google API
type APIKey struct {
	Key string `json:"key"`
}

// Address data for url creatation
type Address struct {
	StreetNumber string
	StreetName   string
	City         string
	State        string
}

// GeoLocation data for url creatation
type GeoLocation struct {
	Latitude  string
	Longitude string
}

// MarshallJSON -
func MarshallJSON(interface{}) {

	// err = json.Unmarshal(byteValue, jsonData)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// UnmarshallJSON -
func UnmarshallJSON(data *[]byte, obj interface{}) {

	err := json.Unmarshal(*data, obj)
	if err != nil {
		log.Fatal(err)
	}

}

// AddressFormatter -
func (g GeoLocation) AddressFormatter(geoloc *GeoLocation, apiKey string) string {
	url := "https://maps.googleapis.com/maps/api/geocode/json?latlng=" + geoloc.Latitude + "," + geoloc.Longitude + "&key=" + apiKey
	return url
}

// AddressFormatter -
func (a Address) AddressFormatter(addr *Address, apiKey string) string {
	var str strings.Builder
	str.WriteString(addr.StreetNumber)
	streetNameSplit := strings.Split(addr.StreetName, " ")
	for i, w := range streetNameSplit {
		if i < len(streetNameSplit) {
			str.WriteString("+" + w)
		}
	}
	str.WriteString(",")
	cityNameSplit := strings.Split(addr.City, " ")
	for i, w := range cityNameSplit {
		if i < len(cityNameSplit) {
			str.WriteString("+" + w)
		}
	}
	str.WriteString(",")
	str.WriteString(addr.State)
	url := "https://maps.googleapis.com/maps/api/geocode/json?address=" + str.String() + "&key=" + apiKey
	return url
}
