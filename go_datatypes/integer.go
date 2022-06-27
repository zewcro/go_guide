// allow to store both positive and negative values
// int	Depends on platform:
// 32 bits in 32 bit systems and
// 64 bit in 64 bit systems	-2147483648 to 2147483647 in 32 bit systems and
// -9223372036854775808 to 9223372036854775807 in 64 bit systems
// int8	8 bits/1 byte	-128 to 127
// int16	16 bits/2 byte	-32768 to 32767
// int32	32 bits/4 byte	-2147483648 to 2147483647
// int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807

package main

import (
	"fmt"
)

func main() {
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
}
