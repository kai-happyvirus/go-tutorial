package main

import (
	"fmt"
)

func vals() (int, int) {
	return 3, 7
}
func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
	variadicFunctions(1, 2)
	variadicFunctions(1, 2, 3)
	nums := []int{1, 2, 3, 4}

	variadicFunctions(nums...)
	nextInt := closures()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := closures()
	fmt.Println(newInts())
}

func variadicFunctions(nums ...int) {
	fmt.Print(nums, ",")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func closures() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
