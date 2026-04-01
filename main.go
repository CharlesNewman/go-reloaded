package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file")
		return
	}

	text := string(data)

	tokens := tokenize(text)
	tokens = applyCommands(tokens)
	tokens = fixArticles(tokens)
	result := buildText(tokens)

	result = result + "\n"

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing output file")
		return
	}
}

func tokenize(text string) []string {
	var tokens []string
	current := ""

	i := 0
	for i < len(text) {
		ch := text[i]

		if ch == ' ' || ch == '\t' {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			i++
			continue
		}

		if ch == '\n' {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, "\n")
			i++
			continue
		}

		if ch == '(' {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}

			command := ""
			command += string(ch)
			i++

			for i < len(text) && text[i] != ')' {
				command += string(text[i])
				i++
			}

			if i < len(text) && text[i] == ')' {
				command += ")"
				i++
			}

			tokens = append(tokens, command)
			continue
		}

		if isPunctuationChar(ch) || ch == '\'' {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}

			if ch == '.' {
				if i+2 < len(text) && text[i+1] == '.' && text[i+2] == '.' {
					tokens = append(tokens, "...")
					i += 3
					continue
				}
			}

			if ch == '!' || ch == '?' {
				group := ""
				for i < len(text) && (text[i] == '!' || text[i] == '?') {
					group += string(text[i])
					i++
				}
				tokens = append(tokens, group)
				continue
			}

			tokens = append(tokens, string(ch))
			i++
			continue
		}

		current += string(ch)
		i++
	}

	if current != "" {
		tokens = append(tokens, current)
	}

	return tokens
}

func applyCommands(tokens []string) []string {
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

func stringToInt(s string) int {
	n := 0

	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		n = n*10 + int(s[i]-'0')
	}

	return n
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

func isPunctuationChar(ch byte) bool {
	if ch == '.' || ch == ',' || ch == '!' || ch == '?' || ch == ':' || ch == ';' {
		return true
	}
	return false
}

func isPunctuationToken(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		if !isPunctuationChar(s[i]) {
			return false
		}
	}

	return true
}

func toUpperManual(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'a' && ch <= 'z' {
			ch = ch - 32
		}
		result += string(ch)
	}

	return result
}

func toLowerManual(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' {
			ch = ch + 32
		}
		result += string(ch)
	}

	return result
}

func capitalizeManual(s string) string {
	if s == "" {
		return s
	}

	s = toLowerManual(s)

	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if i == 0 && ch >= 'a' && ch <= 'z' {
			ch = ch - 32
		}
		result += string(ch)
	}

	return result
}

func hexToDecimalString(s string) string {
	value := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]
		digit := 0

		if ch >= '0' && ch <= '9' {
			digit = int(ch - '0')
		} else if ch >= 'A' && ch <= 'F' {
			digit = int(ch-'A') + 10
		} else if ch >= 'a' && ch <= 'f' {
			digit = int(ch-'a') + 10
		} else {
			return s
		}

		value = value*16 + digit
	}

	return intToString(value)
}

func binToDecimalString(s string) string {
	value := 0

	for i := 0; i < len(s); i++ {
		if s[i] != '0' && s[i] != '1' {
			return s
		}
		value = value*2 + int(s[i]-'0')
	}

	return intToString(value)
}

func intToString(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""

	for n > 0 {
		digit := n % 10
		result = string(byte(digit+'0')) + result
		n = n / 10
	}

	return result
}

func fixArticles(tokens []string) []string {
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

func buildText(tokens []string) string {
	result := ""
	inQuote := false

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token == "'" {
			if !inQuote {
				if result != "" && result[len(result)-1] != ' ' && result[len(result)-1] != '\n' {
					result += " "
				}
				result += "'"
				inQuote = true
			} else {
				result += "'"
				inQuote = false
			}
			continue
		}

		if token == "\n" {
			result = removeLastSpace(result)
			result += "\n"
			continue
		}

		if isPunctuationToken(token) {
			result = removeLastSpace(result)
			result += token
			continue
		}

		if result != "" {
			last := result[len(result)-1]
			if last != '\'' && last != '\n' {
				result += " "
			}
		}

		result += token
	}

	return result
}

func removeLastSpace(s string) string {
	for len(s) > 0 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s
}
