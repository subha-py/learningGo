package main

import (
	"fmt"
)

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Println(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Println(i, " ")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Println(n, " ")
		}(i)
	}
}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
