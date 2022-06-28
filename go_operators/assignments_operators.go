// Assignments Operators are used to assign values to variables

// Exemple :

package main

import (
	"fmt"
)

func main() {
	// assign the value 10 to variable x
	var x = 10
	fmt.Println(x)

	// addition assignment operator (+=) adds value to a variable
	var y = 100
	y += 100
	fmt.Println(y)

}


// A list of all assignment operators:
/*
Operator	Example				Same As	
=				x = 5			x = 5	
+=				x += 3			x = x + 3	
-=				x -= 3			x = x - 3	
*=				x *= 3			x = x * 3	
/=				x /= 3			x = x / 3	
%=				x %= 3			x = x % 3	
&=				x &= 3			x = x & 3	
|=				x |= 3			x = x | 3	
^=				x ^= 3			x = x ^ 3	
>>=				x >>= 3			x = x >> 3	
<<=				x <<= 3			x = x << 3 */