package funcs

func capitalizeManual(s string) string {
	if s == "" {
		return s
	}

	s = toLowerManual(s)

	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if i == 0 && ch >= 'a' && ch <= 'z' {
			ch = ch - 32
		}
		result += string(ch)
	}

	return result
}
