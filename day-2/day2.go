package main

import "fmt"

func main() {
	forLoop()
	ifElse()
	switchStatement()
}

func switchStatement() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
		// Multiple expressions in the same case statement
	case 2, 3, 4:
		fmt.Println("two")
	default:
		fmt.Println("number")
	}

	// Switch without an expression is an alternate way to express if/else logic.
	// Here we also show how the case expressions can be non-constants.
	// This is a clean way to write long if-then-else chains.
	switch {
	case i == 1:
		fmt.Println("one")
	case i == 2:
		fmt.Println("two")
	default:
		fmt.Println("number")
	}
}
func forLoop() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {

		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func ifElse() {

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
