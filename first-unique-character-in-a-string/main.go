func firstUniqChar(s string) int {
	if len(s) < 1 || len(s) > int(math.Pow10(5)) {
		return -1
	}

	m := map[int32]int{}

	for _, char := range s {
		m[char] += 1
	}

	for i := 0; i < len(s); i++ {
		if count, ok := m[int32(s[i])]; ok {
			fmt.Println(i, count)
			if count == 1 {
				return i
			}
		}
	}

	return -1
}