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
