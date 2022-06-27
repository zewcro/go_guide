// Arrays are used to store multiple values of the same type in a single variable,
// instead of declaring separate variables for each value.

package main

import (
	"fmt"
)

func main() {

	//  Array Declaration :

	// with a known size of array

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	// non-known size of array
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8}

	// Print Arrays Informations
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	// Access Elements of an Array

	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Change Elements of an Array

	prices2 := [3]int{10, 20, 30}

	prices2[2] = 50
	fmt.Println(prices2)

	// Find the Length of an Array
	// The len() function is used to find the length of an array:

	arrS := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arrS2 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arrS))
	fmt.Println(len(arrS2))

	// Create an array, named cars of type string

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars)

}
