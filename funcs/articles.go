package funcs

func FixArticles(tokens []string) []string {
	for i := 0; i < len(tokens)-1; i++ {
		if tokens[i] == "a" || tokens[i] == "A" || tokens[i] == "an" || tokens[i] == "An" {
			next := nextWord(tokens, i+1)
			if next != -1 {
				if startsWithVowelOrH(tokens[next]) {
					if tokens[i] == "A" || tokens[i] == "An" {
						tokens[i] = "An"
					} else {
						tokens[i] = "an"
					}
				} else {
					if tokens[i] == "A" || tokens[i] == "An" {
						tokens[i] = "A"
					} else {
						tokens[i] = "a"
					}
				}
			}
		}
	}

	return tokens
}

func nextWord(tokens []string, start int) int {
	for i := start; i < len(tokens); i++ {
		if isWord(tokens[i]) {
			return i
		}
	}
	return -1
}

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
