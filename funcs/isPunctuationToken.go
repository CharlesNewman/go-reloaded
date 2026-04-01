package funcs

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
