//Kieran O'Halloran

package main

//import math and ftm
import (
	"fmt"
	"math"
)

const Delta = 0.0000001
const InitialZ = 600.0

func NewtonsSqrt(x float64) (z float64) {
	const IniitialZ = 100.0
	z = InitialZ

	step := func() float64 {
		return z - (z*z-x)/(2*z)
	}
	for zz := step(); math.Abs(zz-z) > Delta; {
		z = zz
		zz = step()
	}
	return
}

func main() {
	fmt.Println("\nNewtons answer is: ", (NewtonsSqrt(1500)))      //Newtons formula
	fmt.Println("\nThe Math package answer is: ", math.Sqrt(1500)) //callled from math package
}
