package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func login() {

}

func createAccount() {

}

func logInMenu() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("(1) Log In\n" +
			"(2) Create Account\n" +
			"(q) Quit\n" +
			"$ ")
		stdin, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading console")
			continue
		}

		userSelect := strings.ToLower(strings.Trim(stdin, "\n"))
		switch userSelect {
		case "q":
			fmt.Println("Good Bye")
			os.Exit(0)
		case "1":
			login()
		case "2":
			createAccount()
		}
	}
}

func main() {

	// Get User Request
	logInMenu()

	// // Read JSON file, apply struct type, and unmarshall
	// apikeyBytes := filehandler.ReadJSONfile("apikey.json")
	// apikeyStruct := new(apistructs.APIKey)
	// apistructs.UnmarshallJSON(&apikeyBytes, apikeyStruct)

	// address := apistructs.Address{
	// 	StreetNumber: "1600",
	// 	StreetName:   "Amphitheatre Parkway",
	// 	City:         "Mountain View",
	// 	State:        "CA",
	// }

	// geoLocation := apistructs.GeoLocation{
	// 	Latitude:  "40.714224",
	// 	Longitude: "-73.961452",
	// }

	// // ADDRESS LOOKUP
	// formatedURL := address.AddressFormatter(&address, apikeyStruct.Key)
	// fmt.Println("---------ADDRESS LOOKUP----------------")
	// //geocodeBytes, _ := apistructs.RequestURL(formatedURL)
	// // stringData := string(*geocodeBytes)
	// // fileName := "geocodelookup.json"
	// // filehandler.PrintToFile(&stringData, fileName)

	// // REVERSE LOOKUP
	// formatedURL = geoLocation.AddressFormatter(&geoLocation, apikeyStruct.Key)
	// fmt.Println("---------REVERSE LOOKUP----------------")
	// reverBytes, _ := apistructs.RequestURL(formatedURL)
	// // stringData = string(*reverBytes)
	// // fileName = "reverselookup.json"
	// // filehandler.PrintToFile(&stringData, fileName)

	// apistructs.UnstructuredJSON(reverBytes)

}
