/*
Exercise: Loops and Functions

As a simple way to play with functions and loops, implement the square root function using Newton's method.

In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:

To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).

Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta). See if that's more or fewer iterations. How close are you to the math.Sqrt?

Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:

z := float64(1)
z := 1.0

*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	a:=float64(1)
	b:= float64(0)
	//fmt.Println(math.Abs(b-a))
	for math.Abs(b-a ) > 0.001{
		b= a
	a = a - ((a*a - x)/2*a)
		//fmt.Println(b-a ,a, b)
	}

	return a
}

func main() {
	fmt.Println(Sqrt(2))
}
