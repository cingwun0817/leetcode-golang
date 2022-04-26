func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return int(0)
	}

	if len(haystack) == 0 && len(needle) != 0 {
		return int(-1)
	}

	var h, n int
	for {
		if h >= len(haystack) {
			break
		}

		fmt.Println(string(haystack[h]), string(needle[n]))

		if haystack[h] == needle[n] {
			h += 1
			n += 1

			if n == len(needle) {
				return h - n
			}
		} else {
			h = h - n + 1
			n = 0
		}
	}

	return int(-1)
}