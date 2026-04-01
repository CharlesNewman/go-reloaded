package funcs

import "strings"

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
