package main

import "fmt"

func NormalMap() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	// Insert key-value pairs in the map
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for key, value := range countryCapitalMap {
		fmt.Println("Capital of", key, "is", value)
	}

	/* test if entry is present in the map or not*/
	capital, ok := countryCapitalMap["United States"]

	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}
}

// Delete Function in map
func DeleteMapFunction() {
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("Original map")

	// Print map
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	// delete an entry
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("Updated map")

	// Print Map
	for key, value := range countryCapitalMap {
		fmt.Println("Capital of", key, "is", value)
	}

}
func main() {
	NormalMap()
	DeleteMapFunction()
}
