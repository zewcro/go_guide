package main

import (
	"fmt"
)

func main() {
	// This example shows how to create slices using the []datatype{values} format:

	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// Create a slice from an array

	arr1 := [6]int{10, 11, 12, 13, 14, 15}

	myslice := arr1[2:4] // Récupères les valeurs dans cet intervale (par rapport au tableau de base)

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	// In the example above myslice is a slice with length 2.
	// It is made from arr1 which is an array with length 6.

	// The slice starts from the second element of the array which has value 12.
	// The slice can grow to the end of the array. This means that the capacity of the slice is 4.

	// If myslice started from element 0, the slice capacity would be 6.

	// Create a slice using the make() function

	// myslice1 := make([]int, 5, 10)
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	// // with omitted capacity
	// myslice2 := make([]int, 5)
	// fmt.Printf("myslice2 = %v\n", myslice2)
	// fmt.Printf("length = %d\n", len(myslice2))
	// fmt.Printf("capacity = %d\n", cap(myslice2))

	// Access Elements of a Slice

	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Change Elements of a Slice
	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Append Element to a Slice

	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// Append one slice to another slice

	// myslice1 := []int{1, 2, 3}
	// myslice2 := []int{4, 5, 6}
	// myslice3 := append(myslice1, myslice2...)

	// fmt.Printf("myslice3=%v\n", myslice3)
	// fmt.Printf("length=%d\n", len(myslice3))
	// fmt.Printf("capacity=%d\n", cap(myslice3))

	// Change the lenght of a slice

	// arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	// myslice1 := arr1[1:5] // Slice array
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	// myslice1 = arr1[1:3] // Change length by re-slicing the array
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	// myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

}
