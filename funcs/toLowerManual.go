package funcs

func toLowerManual(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' {
			ch = ch + 32
		}
		result += string(ch)
	}

	return result
}
