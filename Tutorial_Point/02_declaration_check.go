package main

import "fmt"

func main(){
	// Static type declaration in go
	var a uint8
	a = 102

	var x = 20
	fmt.Printf("Hello world %d, %d\n", a, x)

	// dynamic type declaration in go
	var temp float32 = 1025.5658484

	y := 20
	fmt.Println(temp)
	fmt.Printf("%d, %T", y, y)

	// Mixed declaration in go
	var aa, bb, cc ,dd = 10, 20.45, 1.0959945, "subhadeep"
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Printf("a is of type %T\n", aa)
	fmt.Printf("b is of type %T\n", bb)
	fmt.Printf("c is of type %T\n", cc)
	fmt.Printf("d is of type %T\n", dd)
}
