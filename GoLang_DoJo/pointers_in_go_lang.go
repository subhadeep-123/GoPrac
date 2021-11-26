package main

import "fmt"

func changeWeapon(existingWeapong *string) {
	*existingWeapong = "Uses a Ninja Sword"
}

func main() {
	fmt.Println("Pointers in go lang")
	ninjaweapon := "Uses a Ninja Star"
	fmt.Println(ninjaweapon)
	changeWeapon(&ninjaweapon)
	fmt.Println(ninjaweapon)
}
