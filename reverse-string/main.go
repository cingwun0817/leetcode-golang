func reverseString(s []byte) {
	size := len(s)
	end := len(s)/2 - 1

	for i := 0; i <= end; i++ {
		x := s[i]
		y := s[size-i-1]

		s[size-i-1] = x
		s[i] = y
	}
}