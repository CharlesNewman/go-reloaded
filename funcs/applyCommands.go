package funcs

func ApplyCommands(tokens []string) []string {
	var result []string

	for i := 0; i < len(tokens); i++ {
		if isCommand(tokens[i]) {
			result = doCommand(result, tokens[i])
		} else {
			result = append(result, tokens[i])
		}
	}

	return result
}
