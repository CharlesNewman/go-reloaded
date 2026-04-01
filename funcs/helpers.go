package funcs

func isPunctuationChar(ch byte) bool {
	if ch == '.' || ch == ',' || ch == '!' || ch == '?' || ch == ':' || ch == ';' {
		return true
	}
	return false
}

func isPunctuationToken(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		if !isPunctuationChar(s[i]) {
			return false
		}
	}

	return true
}
