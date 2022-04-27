func plusOne(digits []int) []int {
	dLen := len(digits)

	// var firstPlus bool
	for i := dLen - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++

			return digits
		}

		digits[i] = 0
	}

	plus := true
	for _, num := range digits {
		if num != 0 && plus {
			plus = false
		}
	}

	if plus {
		digits = append(digits, 0)
		digits[0] = 1
	}

	return digits
}