package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return nil
}

func twoSumHashTable(nums []int, target int) []int {
	numIndex := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := numIndex[complement]; found {
			return []int{index, i}
		}
		numIndex[num] = i
	}

	return nil
}

func main() {
	nums := []int{0, 4, 3, 0}
	target := 0

	sol := twoSum(nums, target)
	for _, value := range sol {
		fmt.Println(value)
	}

	result := twoSumHashTable(nums, target)
	for _, value := range result {
		fmt.Println(value)
	}

}
