package main

import (
	"fmt"
	"math"
)

func varibles() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "shorthand syntax"
	fmt.Println(f)

	const s string = "constant"
	fmt.Println(s)

	const n = 500000000

	const di = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(di))
	fmt.Println(math.Sin(n))

}
