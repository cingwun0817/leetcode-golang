func moveZeroes(nums []int) {
	var pointer int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != pointer {
				nums[pointer] = nums[i]
				nums[i] = 0
			}

			pointer += 1
		}
	}
}