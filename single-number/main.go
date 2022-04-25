func singleNumber(nums []int) int {
	var res int
	size := len(nums)

	for i := 0; i < size; i++ {
		find := false

		for j := 0; j < size; j++ {
			if i != j && nums[i] == nums[j] && !find {
				find = true
			}
		}

		fmt.Println(find, nums[i])

		if !find {
			res = nums[i]
		}
	}

	return res
}