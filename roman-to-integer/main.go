func romanToInt(s string) int {
	var res int
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i, char := range s {
		var num, nextNum int

		num = m[string(char)]

		if i+1 < len(s) {
			nextNum = m[string(s[i+1])]
		}

		if nextNum > num && i+1 < len(s) {
			res -= num
		} else {
			res += num
		}
	}

	return res
}