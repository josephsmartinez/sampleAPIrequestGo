package main

import (
	"bufio"
	"fmt"
	"os"
	"sampleAPIrequestGo/auth"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Start Services i.e. Database
func init() {
	fmt.Println("Database sever started...")
	// Add Docker Command and init start up process
}

// Log User Into Application
func login() {
}

// Create User Account
func createAccount() {

	fmt.Println("Let's create you user profile...")
	for {

		questions := map[string]string{
			"firstName":    "Please enter your first name",
			"lastName":     "Please enter your last name",
			"password":     "Please enter a password",
			"apikey":       "enter you api key",
			"streetnumber": "Please enter your street number",
			"streetname":   "Please enter your street name",
			"city":         "Please enter your city code",
			"state":        "Please enter your state",
		}

		fmt.Printf("%v\n$ ", questions["firstName"])
		firstName, err := reader.ReadString('\n')
		fmt.Printf("%v\n$ ", questions["lastName"])
		lastName, err := reader.ReadString('\n')
		fmt.Printf("%v\n$ ", questions["password"])
		password, err := reader.ReadString('\n')
		fmt.Printf("%v\n$ ", questions["apikey"])
		apikey, err := reader.ReadString('\n')
		fmt.Printf("%v\n$ ", questions["streetnumber"])
		streetnumber, err := reader.ReadString('\n')
		fmt.Printf("%v\n$ ", questions["streetname"])
		streetname, err := reader.ReadString('\n')
		fmt.Printf("%v\n$ ", questions["city"])
		city, err := reader.ReadString('\n')
		fmt.Printf("%v\n$ ", questions["state"])
		state, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading console")
			continue
		}
		newUser := auth.NewUser()
		newUser.FirstName = strings.Trim(firstName, "\n")
		newUser.LastName = strings.Trim(lastName, "\n")
		newUser.Password = strings.Trim(password, "\n")
		newUser.APIKey = strings.Trim(apikey, "\n")
		newUser.StreetNumber = strings.Trim(streetnumber, "\n")
		newUser.StreetName = strings.Trim(streetname, "\n")
		newUser.State = strings.Trim(state, "\n")
		newUser.City = strings.Trim(city, "\n")

		// Check if user information us correct
		fmt.Println(*newUser)
		fmt.Println("Is the following information correct? [y/n]")
		corrInfo, err := reader.ReadString('\n')
		corrInfo = strings.ToLower(strings.Trim(corrInfo, "\n"))
		if corrInfo == "y" {
			break
		}
	}

}

func logInMenu() {
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

	// Startup Services: Docker SQL

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
