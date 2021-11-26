package main

import "fmt"

type ninja struct {
	name    string
	weapons []string
	level   int
}

type SoftwareNinjs struct {
	company string
	info    ninja
}

type SoftwareNinja struct {
	company string
	details *ninja
}

func testArray() {
	var arr []int
	fmt.Println(arr)
	arr = append(arr, 5, 6, 8, 7)
	fmt.Println(arr)
}

func dynamicStruct() {
	dstruct := new(ninja)
	dstruct.name = "Subhadeep Banerjee"
	dstruct.level = 3
	dstruct.weapons = []string{"heart", "soul", "will power"}
	fmt.Println(dstruct)
}
func main() {
	matrix := ninja{
		name:    "Subhadeep Banerje",
		weapons: []string{"python", "go"},
		level:   3}
	fmt.Println(matrix)
	fmt.Println(matrix.name)
	fmt.Println(matrix.weapons)
	matrix.weapons = append(matrix.weapons, "Bash", "HTML", "CSS")
	fmt.Println(matrix.level)
	fmt.Println(matrix)
	// testArray()

	sn := SoftwareNinjs{company: "Klizo Solutions", info: matrix}
	sn2 := SoftwareNinja{company: "Klizo Solutions", details: &matrix}
	fmt.Println(sn)
	fmt.Println(*sn2.details)
	matrix.weapons = append(matrix.weapons, "Docker", "AWS", "Git")
	// fmt.Println(matrix)
	fmt.Println(sn)
	fmt.Println(*sn2.details)
	dynamicStruct()
}
