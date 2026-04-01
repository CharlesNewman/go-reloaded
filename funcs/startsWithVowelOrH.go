package funcs

func startsWithVowelOrH(word string) bool {
	if word == "" {
		return false
	}

	first := word[0]

	if first >= 'A' && first <= 'Z' {
		first = first + 32
	}

	if first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' {
		return true
	}

	return false
}
