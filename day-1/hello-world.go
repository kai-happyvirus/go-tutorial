package main

import "fmt"

func variable() {

	// equivalent var short string = "short"
	// Only available inside a function
	short := "short"
	fmt.Println(short)

	var num int = 1
	num1 := 2
	fmt.Println(num)
	fmt.Println(num1)

}

func typeValue() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func main() {
	fmt.Println("Hello, World!")
	typeValue()
	variable()
}
