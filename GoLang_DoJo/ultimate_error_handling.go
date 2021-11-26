package main

import (
	"errors"
	"fmt"
	"strconv"
)

func returnError(returnError bool) (string, error) {
	if returnError {
		return "", errors.New("Some Error Here")
	}
	return "Here is the output", nil
}


type SpecialCustomError struct {
	errorMessage string
	errorCode    int
}

func (s SpecialCustomError) Error() string {
	return s.errorMessage + " " + strconv.Itoa(s.errorCode)
}

func returnSpecialError(returnSpecialError bool) (string, error) {
	if returnSpecialError {
		return "", SpecialCustomError{"Some Error Here", 123}
	}
	return "Here is the output", nil
}

func main() {
	val, err := returnError(true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}


	spec_val, spec_err := returnSpecialError(true)
	if spec_err != nil {
		fmt.Println(spec_err)
	} else {
		fmt.Println(spec_val)
	}
}
