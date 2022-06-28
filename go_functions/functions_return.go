package main

import (
	"fmt"
)

func myFunction(x int, y int) int {
	return x + y
}

func myFunction2(x int, y int) (result int) {
	result = x + y
	return result
}

func main() {
	fmt.Println(myFunction(1, 2))
	fmt.Println(myFunction(10000, 678393))
	fmt.Println(myFunction2(20, 240))

	somme := myFunction2(10, 10)
	fmt.Println(somme)
}
