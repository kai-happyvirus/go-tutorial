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
}

func variadicFunctions(nums ...int) {
	fmt.Print(nums, ",")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
