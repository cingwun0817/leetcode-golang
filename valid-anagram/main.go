func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[string]int{}

	for _, char := range s {
		_, ok := m[string(char)]

		if !ok {
			m[string(char)] = 1
		} else {
			m[string(char)] += 1
		}
	}

	for _, char := range t {
		val, ok := m[string(char)]

		if !ok || val == 0 {
			return false
		}

		m[string(char)] -= 1
	}

	return true
}