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
