package funcs

import "strings"

func findCommands(tokens []string) bool {

	isCommand := false

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(hex)" || tokens[i] == "(bin)" || tokens[i] == "(up)" || tokens[i] == "(low)" || tokens[i] == "(cap)" {
			isCommand = true
		}
		if tokens[i] == "(up," || tokens[i] == "(low," || tokens[i] == "(cap," {
			if i+1 < len(tokens) {
				if strings.HasSuffix(tokens[i+1], ")") {
					isCommand = true
				}
			}
		}
	}
	return isCommand
}
