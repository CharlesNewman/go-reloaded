package funcs

func FixArticles(tokens []string) []string {
	for i := 0; i < len(tokens)-1; i++ {
		if tokens[i] == "a" || tokens[i] == "A" || tokens[i] == "an" || tokens[i] == "An" {
			next := nextWord(tokens, i+1)
			if next != -1 {
				if startsWithAnSound(tokens[next]) {
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

func startsWithAnSound(word string) bool {
	if word == "" {
		return false
	}

	w := toLowerManual(word)

	for len(w) > 0 && (w[0] == '\'' || w[0] == '"' || isPunctuationChar(w[0])) {
		w = w[1:]
	}

	if w == "" {
		return false
	}

	if hasYouSound(w) {
		return false
	}

	if hasSilentHPrefix(w) {
		return true
	}

	first := w[0]
	return first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u'
}
