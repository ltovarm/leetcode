package majority

func majorityElement(nums []int) int {
	counts := make(map[int]int)
	result := 0

	for _, value := range nums {
		counts[value]++
		if counts[value] > counts[result] {
			result = value
		}
	}
	return result
}
