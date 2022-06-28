package main

import (
	"fmt"
)

// Function with One Parameter
func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

// Function with Multiple Parameters
func familyNames(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Doe")
}

func main() {
	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")
	familyNames("Theo", 22)
	familyNames("Bob", 20)
	familyNames("John", 100)
}
