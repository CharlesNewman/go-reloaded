package funcs

func doCommand(tokens []string, command string) []string {
	name, count := readCommand(command)

	if name == "" {
		return tokens
	}

	if name == "hex" {
		index := findPreviousWord(tokens)
		if index != -1 {
			tokens[index] = hexToDecimalString(tokens[index])
		}
		return tokens
	}

	if name == "bin" {
		index := findPreviousWord(tokens)
		if index != -1 {
			tokens[index] = binToDecimalString(tokens[index])
		}
		return tokens
	}

	if name == "up" || name == "low" || name == "cap" {
		done := 0

		for i := len(tokens) - 1; i >= 0; i-- {
			if isWord(tokens[i]) {
				if name == "up" {
					tokens[i] = toUpperManual(tokens[i])
				}
				if name == "low" {
					tokens[i] = toLowerManual(tokens[i])
				}
				if name == "cap" {
					tokens[i] = capitalizeManual(tokens[i])
				}

				done++
				if done == count {
					break
				}
			}
		}
	}

	return tokens
}
