package main

import "fmt"

// Variadic function : it has the ability to accept a variable number of arguments. Eg fmt.Println(1, "hello", true)
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total
}

func print() {
	// result := sum(1, 2, 3, 4, 5)
	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println("Sum:", result)
}
