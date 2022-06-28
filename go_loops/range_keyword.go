package main

import (
	"fmt"
)

func main() {

	fruits := [3]string{"apple", "orange", "banana"}
	for i, v := range fruits {
		fmt.Printf("%v\t%v\n", i, v)
	}

	legumes := [3]string{"spinach", "brocoli", "potatoe"}
	for _, legume := range legumes {
		fmt.Printf("%v\n", legume)
	}
}
