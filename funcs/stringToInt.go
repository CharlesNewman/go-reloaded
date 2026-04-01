package funcs

func stringToInt(s string) int {
	n := 0

	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		n = n*10 + int(s[i]-'0')
	}

	return n
}
