package rotate_array

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate(nums []int, k int) {

	lenght := len(nums)

	// Adjust n in case k > lenght avoid incessant rotations
	k = k % lenght

	reverse(nums, 0, lenght-1) // Reverse entire array
	reverse(nums, 0, k-1)      // Reverse first half of array
	reverse(nums, k, lenght-1) // Reverse second half of array
}

// func main() {

// 	nums := []int{1, 2, 3, 4, 5, 6, 7}
// 	k := 3

// 	rotate(nums, k)

// 	fmt.Println(nums)

// 	nums = []int{-1, -100, 3, 99}
// 	k = 2
// 	rotate(nums, k)

// 	fmt.Println(nums)

// 	nums = []int{-1}
// 	k = 2
// 	rotate(nums, k)

// 	fmt.Println(nums)

// 	nums = []int{1, 2}
// 	k = 3
// 	rotate(nums, k)

// 	fmt.Println(nums)

// 	nums = []int{1, 2, 3, 4, 5, 6, 7}
// 	k = 9

// 	rotate(nums, k)

// 	fmt.Println(nums)

// }
