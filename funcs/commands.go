package funcs

import "strings"

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

func readCommand(command string) (string, int) {
	if len(command) < 2 {
		return "", 1
	}

	if command[0] != '(' || command[len(command)-1] != ')' {
		return "", 1
	}

	inside := command[1 : len(command)-1]

	if inside == "up" {
		return "up", 1
	}
	if inside == "low" {
		return "low", 1
	}
	if inside == "cap" {
		return "cap", 1
	}
	if inside == "hex" {
		return "hex", 1
	}
	if inside == "bin" {
		return "bin", 1
	}

	parts := strings.Split(inside, ",")

	if len(parts) != 2 {
		return "", 1
	}

	name := strings.TrimSpace(parts[0])
	numberText := strings.TrimSpace(parts[1])

	count := stringToInt(numberText)
	if count <= 0 {
		count = 1
	}

	if name == "up" || name == "low" || name == "cap" {
		return name, count
	}

	return "", 1
}

func findPreviousWord(tokens []string) int {
	for i := len(tokens) - 1; i >= 0; i-- {
		if isWord(tokens[i]) {
			return i
		}
	}
	return -1
}

func isCommand(s string) bool {
	if len(s) < 2 {
		return false
	}
	return s[0] == '(' && s[len(s)-1] == ')'
}

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
