/*


Operator	Name					Description															Example
& 			AND						Sets each bit to 1 if both bits are 1								x & y
|			OR						Sets each bit to 1 if one of two bits is 1							x | y
^			XOR						Sets each bit to 1 if only one of two bits is 1						x ^ b
<<			Zero 					fill left shift		Shift left by pushing zeros in from the right	x << 2
>>			Signed right shift		Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off	x >> 2


*/

// Exemple for XOR

package main

import (
	"fmt"
)

func main() {

	var x = 9
	var y = 8

	fmt.Printf("x = %b\n", x)
	fmt.Printf("y = %b\n", y)

	fmt.Printf("x ^ y is %b\n", x^y)
}
