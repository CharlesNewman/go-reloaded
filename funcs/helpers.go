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

func hasSilentHPrefix(word string) bool {
	silentHWords := []string{
		"heir",
		"heiress",
		"honest",
		"honesty",
		"honor",
		"honour",
		"honorable",
		"honourable",
		"hour",
		"hourly",
		"hourglass",
	}

	for i := 0; i < len(silentHWords); i++ {
		prefix := silentHWords[i]
		if len(word) >= len(prefix) && word[:len(prefix)] == prefix {
			return true
		}
	}

	return false
}

func hasYouSound(word string) bool {
	youSoundWords := []string{
		"uni",
		"use",
		"user",
		"usual",
		"utensil",
		"utility",
		"euro",
		"europe",
		"european",
		"one",
		"once",
	}

	for i := 0; i < len(youSoundWords); i++ {
		prefix := youSoundWords[i]
		if len(word) >= len(prefix) && word[:len(prefix)] == prefix {
			return true
		}
	}

	return false
}
