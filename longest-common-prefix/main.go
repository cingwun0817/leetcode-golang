func longestCommonPrefix(strs []string) string {
	var res string

	for i := 0; i < 200; i++ {
		var firstChar byte
		isSame := true

		for index, str := range strs {
			var char byte

			if i < len(str) {
				char = str[i]
			}

			if index == 0 {
				firstChar = char
			}

			if isSame {
				isSame = firstChar == char

				if firstChar == 0 && char == 0 {
					isSame = false
				}
			}

			fmt.Println(firstChar, char, isSame)
		}

		fmt.Println("--- res: ", res)

		if isSame {
			res += string(firstChar)
		} else {
			return res
		}
	}

	return res
}