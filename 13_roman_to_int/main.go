package main

import "fmt"

func romanToInt(s string) int {
	hash := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0
	k := hash[string(s[0])]
	for i := 1; i < len(s); i++ {
		j, _ := hash[string(s[i])]
		if k < j {
			result -= k
		} else {
			result += k
		}
		k = j
	}

	result += k
	return result
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
