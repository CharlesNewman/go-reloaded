package funcs

func toUpperManual(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'a' && ch <= 'z' {
			ch = ch - 32
		}
		result += string(ch)
	}

	return result
}
