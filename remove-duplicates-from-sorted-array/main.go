func removeDuplicates(nums []int) int {
	nowPos := 0
	m := map[int]bool{}
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			nums[nowPos] = num

			nowPos += 1
			m[num] = true
		}
	}

	return nowPos
}