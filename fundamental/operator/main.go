package main

import (
	"fmt"
)

var x int = 9
var y int = 12

func main() {

	fmt.Println(x*y, x+y, x-y, x/y, x%y)

	/*
		Operator	Name				Description	Example
		+			Addition			Adds together two values	x + y
		-			Subtraction			Subtracts one value from another	x - y
		*			Multiplication		Multiplies two values	x * y
		/			Division			Divides one value by another	x / y
		%			Modulus				Returns the division remainder	x % y
		++			Increment			Increases the value of a variable by 1	x++
		--			Decrement			Decreases the value of a variable by 1	x--
	*/

	x += y
	fmt.Println(x)
	x *= y
	fmt.Println(x)
	x /= y
	fmt.Println(x)
	x -= y
	fmt.Println(x)
	x %= y
	fmt.Println(x)

	/*
		Operator	Example		Same As
		=			x = 5		x = 5
		+=			x += 3		x = x + 3
		-=			x -= 3		x = x - 3
		*=			x *= 3		x = x * 3
		/=			x /= 3		x = x / 3
		%=			x %= 3		x = x % 3
	*/

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)

	/*
	   Operator		Name							Example
	   ==			Equal to						x ==  y
	   !=			Not equal						x !=  y
	   >			Greater than					x >	  y
	   <			Less than						x <   y
	   >=			Greater than or equal to		x >=  y
	   <=			Less than or equal to			x <=  y
	*/

	a := true
	b := false
	fmt.Println("Logical operators:")
	fmt.Println(!a)     // Logical NOT (!): returns the opposite boolean
	fmt.Println(a && b) // Logical AND (&&): if both are
	fmt.Println(a || b) // Logical OR (||): if either

}
