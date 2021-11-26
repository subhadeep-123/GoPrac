package main

import "fmt"

type ninjastars struct {
	owner string
}

func (nst ninjastars) attack() {
	fmt.Printf("Throwing Ninja Stars - %s\n", nst.owner)
}

type ninjasord struct {
	owner string
}

func (ns ninjasord) attack() {
	fmt.Printf("Throwing Ninja Swords - %s\n", ns.owner)
}

type ninjaweapon interface {
	attack()
}

func attack(weapon ninjaweapon) {
	weapon.attack()
}
func main() {
	stars := []ninjastars{{"Subhadeep"}, {"Matrix"}}
	for _, star := range stars {
		star.attack()
	}
	fmt.Println()

	swords := []ninjasord{{"Subhadeep"}, {"Matrix"}}
	for _, sword := range swords {
		sword.attack()
	}
	fmt.Println()

	weapons := []ninjaweapon{ninjastars{"Subhadeep"}, ninjasord{"Matrix"}}
	for _, weapon := range weapons {
		weapon.attack()
	}
	fmt.Println()

}
