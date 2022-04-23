// O(n^2)
/*func isAnagrams(p string, t string) bool {
    if len(p) != len(t) {
        return false
    }

    m := map[string]int{}
    for _, char := range p {
        _, ok := m[string(char)]

        if !ok {
            m[string(char)] = 1
        } else {
            m[string(char)] += 1
        }
    }

    for _, char := range t {
        key, ok := m[string(char)]

        if !ok || key == 0 {
            return false
        }

        m[string(char)] -= 1
    }

    return true
}*/

func findAnagrams(s string, p string) []int {
	// O(n)
	res := []int{}

	if len(s) == 0 || len(s) < len(p) {
		return res
	}

	var vp [26]int
	for i := 0; i < len(p); i++ {
		vp[p[i]-'a']++
	}

	var vs [26]int
	for i := 0; i < len(s); i++ {
		if i >= len(p) {
			vs[s[i-len(p)]-'a']--
		}

		vs[s[i]-'a']++

		if vs == vp {
			res = append(res, i-len(p)+1)
		}
	}

	return res

	/*res := []int{}

	    fmt.Println(len(s), len(p))
	    if len(s) == 0 || len(s) < len(p) {
			return res
		}

	    pLen := len(p)

	    for i, _ := range s {
	        if (len(s) - i) < pLen {
	            break
	        }

	        chars := s[i:i+pLen]
	        ok := isAnagrams(p, chars)

	        if ok {
	            res = append(res, i)
	        }
	    }

	    return res*/
}