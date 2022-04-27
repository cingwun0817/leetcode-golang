func missingNumber(nums []int) int {
	numLen := len(nums)

	var x, y int
	for i := 0; i < numLen; i++ {
		x += i + 1
		y += nums[i]
	}

	fmt.Println(x, y)

	return x - y

	return 0
}