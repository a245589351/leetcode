package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	another := 0

	for i := range nums {
		another = target - nums[i]
		if j, ok := hash[another]; ok {
			return []int{i, j}
		}
		hash[nums[i]] = i
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)

	nums = []int{3, 2, 4}
	target = 6
	result = twoSum(nums, target)
	fmt.Println(result)

	nums = []int{3, 3}
	target = 6
	result = twoSum(nums, target)
	fmt.Println(result)
}
