package main

import (
	"fmt"
)

func myFunction3(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {
	fmt.Println(myFunction3(5, "Hello"))
}
