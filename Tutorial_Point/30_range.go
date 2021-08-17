package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Print the numbers
	for i, j := range numbers {
		fmt.Printf("Array Index - %d, Value - %d\n", i, j)
	}

	// for i := range numbers {
	// 	fmt.Printf("Array Value - %d\n", numbers[i])
	// }

	// Creating a map
	countryCapitalMap := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo"}

	for k, v := range countryCapitalMap {
		fmt.Printf("Key - %s, Value - %s\n", k, v)
	}

}
