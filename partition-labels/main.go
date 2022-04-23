func partitionLabels(s string) []int {
	var res []int

	pos := make(map[int32][]int)
	for i, c := range s {
		pos[c] = append(pos[c], i)
	}

	var start, end int
	for i, c := range s {
		charEnd := pos[c][len(pos[c])-1]
		if charEnd > end {
			end = charEnd
		}

		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}

		//fmt.Println(i, string(c), end)
	}

	return res
}