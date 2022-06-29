package main

import (
	"fmt"
)

type Personne struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Personne
	var pers2 Personne

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Print Pers1 info by calling a function
	printPerson(pers1)

	// Print Pers2 info by calling a function
	printPerson(pers2)
}

func printPerson(pers Personne) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
