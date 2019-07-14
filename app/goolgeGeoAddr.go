package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sampleAPIrequestGo/apistructs"
	"sampleAPIrequestGo/filehandler"
)

// Global Structs
var reader = bufio.NewReader(os.Stdin)

// Read JSON file, apply struct type, and unmarshall
var apikeyBytes []byte = filehandler.ReadJSONfile("jsondata/apikey.json")
var apikeyStruct = new(apistructs.APIKey)

// Simple User Data
var userdata = apistructs.User{
	FirstName: "John",
	LastName:  "Doe",
	Password:  "pass123",
	Address:   "1600 Amphitheatre Parkway",
	Address2:  "",
	City:      "Mountain View",
	State:     "CA",
}

func googleGeocodingAPI() {

	address := apistructs.Address{
		StreetNumber: "1600",
		StreetName:   "Amphitheatre Parkway",
		City:         "Mountain View",
		State:        "CA",
	}

	geoLocation := apistructs.GeoLocation{
		Latitude:  "40.714224",
		Longitude: "-73.961452",
	}

	// ADDRESS LOOKUP
	formatedURL := address.AddressFormatter(&address, apikeyStruct.Key)
	fmt.Println("---------ADDRESS LOOKUP----------------")
	geocodeBytes, _ := apistructs.RequestURL(formatedURL)
	stringData := string(*geocodeBytes)
	fileName := "jsondata/geocodelookup.json"
	filehandler.PrintToFile(&stringData, fileName)

	// REVERSE LOOKUP
	formatedURL = geoLocation.AddressFormatter(&geoLocation, apikeyStruct.Key)
	fmt.Println("---------REVERSE LOOKUP----------------")
	reverBytes, _ := apistructs.RequestURL(formatedURL)
	stringData = string(*reverBytes)
	fileName = "jsondata/reverselookup.json"
	filehandler.PrintToFile(&stringData, fileName)

	// Just prints data into the console as an example
	apistructs.UnstructuredJSON(reverBytes)
}

func directionsAPI() {

	req := map[string]string{
		"origin":      "NewYork",
		"destination": "Washington",
	}

	url := "https://maps.googleapis.com/maps/api/directions/json?origin=" +
		req["origin"] + "&destination=" + req["destination"] +
		"&key=" + apikeyStruct.Key

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringData := string(body)
	fileName := "jsondata/directions.json"
	filehandler.PrintToFile(&stringData, fileName)

}

func main() {

	// Read JSON file, apply struct type, and unmarshall
	apistructs.UnmarshallJSON(&apikeyBytes, apikeyStruct)

	googleGeocodingAPI()
	directionsAPI()

}
