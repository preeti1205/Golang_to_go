/*
Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer:

type error interface {
   Error() string
}

(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

i, err := strconv.Atoi("42")
if err != nil {
   fmt.Printf("couldn't convert number: %v\n", err)
   return
}
fmt.Println("Converted integer:", i)

A nil error denotes success; a non-nil error denotes failure.

*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {

	if(x < 0){
		return 0, ErrorHandler(x)
	}
	a:=float64(1)
	b:= float64(0)

	for math.Abs(b-a ) > 0.001{
		b= a
	a = a - ((a*a - x)/2*a)
	}

	return a, nil
}

type ErrNegativeSqrt float64

func ErrorHandler(num float64) error{
	return ErrNegativeSqrt(num)
}

func (e ErrNegativeSqrt) Error() string{
	//Sprintf() results in creating a string without printing
	//use Sprint() if you just want to concatentate without formatting (no %)
	return fmt.Sprintf("math: cannot Sqrt Negative number: %g",float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
