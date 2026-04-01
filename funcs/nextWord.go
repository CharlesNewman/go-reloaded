package funcs

func nextWord(tokens []string, start int) int {
	for i := start; i < len(tokens); i++ {
		if isWord(tokens[i]) {
			return i
		}
	}
	return -1
}
