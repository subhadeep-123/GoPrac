package main

import "fmt"

func sliceArray() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println("The Length of the array is - ", len(arr))
	fmt.Println("Sliced Array - ", arr[:2])
	var temp []int = arr[1:5]
	fmt.Println("Sliced Array - ", temp)
}
func sliceString() {
	names := []string{
		"Subhadeep",
		"Saanvi",
		"Dhan",
		"David",
		"Tengku",
		"Naimul"}
	fmt.Println(names)
	s1 := names[0:2]
	s2 := names[1:3]
	fmt.Println(s1, s2)
	s2[0] = "Saanvi Mogla"
	fmt.Println(s1, s2)
	fmt.Println(names)
}
func OmitLowerUpper() {
	arr := []int{1, 2, 3, 4, 5}
	s1 := arr[2:4]
	printSlice(s1)
	s2 := arr[:3]
	printSlice(s2)
	s3 := arr[:2]
	printSlice(s3)
}
func printSlice(s []int) {
	fmt.Printf("length =%d capacity=%d %v\n", len(s), cap(s), s)
}
func usingMakeFunc() {
	// arr := []int{1, 2, 3, 4, 5}
	slice := make([]int, 10)
	printSlice(slice)
	slice1 := make([]int, 1, 3)
	printSlice(slice1)
}
func main() {
	sliceArray()
	sliceString()
	OmitLowerUpper()
	usingMakeFunc()
}
