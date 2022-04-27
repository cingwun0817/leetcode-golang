func isHappy(n int) bool {
	m := map[int]bool{}
	for i := 0; i < 100; i++ {
		n = happyNumber(n)

		if _, ok := m[n]; ok {
			return false
		}

		m[n] = true

		fmt.Println(n)

		if n == 1 {
			return true
		}
	}

	return false
}

func happyNumber(n int) int {
	var res int
	for {
		x := n % 10
		res += int(math.Pow(float64(x), 2))
		n = n / 10

		if n == 0 {
			break
		}
	}

	return res
}