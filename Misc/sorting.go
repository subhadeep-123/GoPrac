package main

import (
	"fmt"
	"sort"
)

type OrderByLengthDesc []string

func (s OrderByLengthDesc) Len() int {
	return len(s)
}
func (str OrderByLengthDesc) Swap(i, j int) {
	str[i], str[j] = str[j], str[i]
}
func (s OrderByLengthDesc) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func main() {
	intValue := []int{10, 20, 5, 3}
	sort.Ints(intValue)
	fmt.Println(intValue)

	floatValue := []float64{100.548, 5.151, 1.0059594594, 0.65695, 65, 5.0612}
	sort.Float64s(floatValue)
	fmt.Println(floatValue)

	name := []string{"Subhadeep", "Richard", "Ria", "Shubham", "Suvhradip"}
	sort.Strings(name)
	fmt.Println(name)

	str := sort.Float64sAreSorted(floatValue)
	fmt.Println(str)

	city := []string{"New York", "London", "Washington", "Delhi"}
	sort.Sort(OrderByLengthDesc(city))
	fmt.Println(city)
}
