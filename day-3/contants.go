package main

import (
	"fmt"
	"math"
)

func main() {
	const a = 50
	fmt.Println(a)

	// declare group of constants
	const (
		name    = "Kai"
		age     = 30
		country = "Germany"
	)

	fmt.Println(name, age, country)

	// The value of a constant should be known at compile time.
	var b = math.Sqrt(4)
	// const c = math.Sqrt(4) // Not allowed
	fmt.Println(b)
	const n = "Sam"
	var nameOfn = n
	fmt.Println("type %T value %v", nameOfn, nameOfn)

	var defaultName = "Sam" //allowed
	type myString string
	var customName myString = "Sam" //allowed
	// customName = defaultName //not allowed
	fmt.Println(customName)
	fmt.Println(defaultName)
	numberConstant()

}

func numberConstant() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}
