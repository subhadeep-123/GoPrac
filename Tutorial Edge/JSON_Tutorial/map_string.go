package main

import (
	"encoding/json"
	"fmt""
)

func main() {
	str := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	var incommingData map[string]interface{}
	err := json.Unmarshal([]byte(str), &incommingData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(incommingData)
}
