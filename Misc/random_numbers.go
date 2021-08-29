package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("%d\n", rand.Intn(100))
	fmt.Printf("%f\n", rand.Float64())
	rand.Seed(time.Now().Unix())
	fmt.Print(rand.Intn(20-1) + 1)
}
