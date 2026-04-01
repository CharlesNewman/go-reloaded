package funcs

func findPreviousWord(tokens []string) int {
	for i := len(tokens) - 1; i >= 0; i-- {
		if isWord(tokens[i]) {
			return i
		}
	}
	return -1
}
