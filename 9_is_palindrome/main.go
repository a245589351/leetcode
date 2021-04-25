package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reInt := 0
	for reInt < x {
		reInt = reInt*10 + x%10
		x /= 10
	}

	return reInt == x || x == reInt/10
}

func main() {
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(-10))
	fmt.Println(isPalindrome(-101))
}
