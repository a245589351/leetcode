package main

import (
	"fmt"
)

func reverse(x int) int {
	result := 0
	maxInt := 1<<31 - 1
	minInt := -1 << 31
	for x != 0 {
		result = result*10 + x%10
		x = x / 10

		if result > maxInt || result < minInt {
			return 0
		}
	}

	return result
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(-120))
	fmt.Println(reverse(-0))
}
