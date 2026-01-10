package main

import "fmt"

//pass by value  [like a copy]
func changeNum(num int) {
	num = 5
	fmt.Println("In ChangeNum", num)
}

//pass by reference
func refChangeNum(num *int) {

	*num = 5
	fmt.Println("In refChangeNum", num)

}

func main() {
	num := 1

	// changeNum(num)
	fmt.Println("Memory Address", &num)

	refChangeNum(&num)
	fmt.Println("AfterChange", num)
	fmt.Println("In main func", num)

	testStructs()
}
