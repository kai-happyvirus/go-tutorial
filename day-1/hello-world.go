package main

import "fmt"

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
}
