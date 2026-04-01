package funcs

func isWord(s string) bool {
	if s == "" {
		return false
	}

	if s == "'" {
		return false
	}

	if s == "\n" {
		return false
	}

	if isCommand(s) {
		return false
	}

	if isPunctuationToken(s) {
		return false
	}

	return true
}
