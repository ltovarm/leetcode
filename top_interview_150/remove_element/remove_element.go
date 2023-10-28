package remove_element

func removeElement(nums []int, val int) int {

	// Initialize values
	result := 0
	i := 0

	// goes through the array and counts the elements that match val
	// and modifies the nums slice.
	for _, value := range nums {
		if value != val {
			result++
			nums[i] = value
			i++
		}
	}

	return result
}
