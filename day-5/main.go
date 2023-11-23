package main

import "fmt"

func main() {

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println(recursion(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursion(n-1)
}
