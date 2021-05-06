package main

import (
	"fmt"
)

func isValid(s string) bool {
	l := len(s)
	if l%2 != 0 {
		return false
	}

	maps := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, 0)
	for i := 0; i < l; i++ {
		if n, ok := maps[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != n {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
