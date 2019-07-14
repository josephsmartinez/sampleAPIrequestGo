package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sampleAPIrequestGo/apistructs"
	"sampleAPIrequestGo/filehandler"
)

// GeocodingRequest HTTP request function
func requestURL(url string) (*[]byte, bool) {
	requestMade := true
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		requestMade = false
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		requestMade = false
	}
	return &body, requestMade
}

func main() {

	// Read JSON file, apply struct type, and unmarshall
	apikeyBytes := filehandler.ReadJSONfile("apikey.json")
	apikeyStruct := new(apistructs.APIKey)
	apistructs.UnmarshallJSON(&apikeyBytes, apikeyStruct)

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
	//geocodeBytes, _ := requestURL(formatedURL)
	// stringData := string(*geocodeBytes)
	// fileName := "geocodelookup.json"
	// filehandler.PrintToFile(&stringData, fileName)

	// REVERSE LOOKUP
	formatedURL = geoLocation.AddressFormatter(&geoLocation, apikeyStruct.Key)
	fmt.Println("---------REVERSE LOOKUP----------------")
	reverBytes, _ := requestURL(formatedURL)
	// stringData = string(*reverBytes)
	// fileName = "reverselookup.json"
	// filehandler.PrintToFile(&stringData, fileName)

	var dat map[string]interface{}
	if err := json.Unmarshal(*reverBytes, &dat); err != nil {
		panic(err)
	}

	statusCode := dat["status"]
	plusCode := dat["plus_code"].(map[string]interface{})
	compoundCode := plusCode["compound_code"]
	globalCode := plusCode["global_code"]

	xresults := dat["results"].([]interface{})
	results := xresults[0].(map[string]interface{})
	xaddressComponents := results["address_components"].([]interface{})
	addressComponents := xaddressComponents[0].(map[string]interface{})

	fmt.Println(statusCode)
	fmt.Println(plusCode)
	fmt.Println(compoundCode)
	fmt.Println(globalCode)
	fmt.Println(addressComponents["long_name"])

}
