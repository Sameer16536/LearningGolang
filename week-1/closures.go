package main

func counter() func() int {
	var count = 0

	return func() int {
		count++
		return count
	}
}

func closures() {
	countFunc := counter()
	println(countFunc()) // Output: 1
}
