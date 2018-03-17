package stringst

func Reverse(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func Kmf(str string) []int {
	// p := []rune(str)
	pi := make([]int, len(str))

	pi[0] = 0
	for k := 1; k < len(str); k++ {
		if str[pi[k-1]] == str[k] {
			pi[k] = pi[k-1] + 1
		} else {
			pi[k] = 0
		}
	}

	return pi
}
