package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong input expected: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	result := processText(string(data))
	result = result + "\n"

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}
}

func processText(text string) string {
	tokens := tokenize(text)
	tokens = applyCommands(tokens)
	tokens = fixArticles(tokens)
	return rebuildText(tokens)
}

func tokenize(text string) []string {
	var tokens []string
	current := ""

	for i := 0; i < len(text); i++ {
		ch := text[i]

		if ch == ' ' || ch == '\n' || ch == '\t' {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			continue
		}

		if ch == '(' {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}

			cmd := ""
			for i < len(text) && text[i] != ')' {
				cmd += string(text[i])
				i++
			}
			if i < len(text) {
				cmd += string(text[i])
			}
			tokens = append(tokens, cmd)
			continue
		}

		if isPunctuationChar(ch) {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}

			if ch == '.' && i+2 < len(text) && text[i+1] == '.' && text[i+2] == '.' {
				tokens = append(tokens, "...")
				i += 2
				continue
			}

			if ch == '\'' {
				tokens = append(tokens, "'")
				continue
			}

			punc := string(ch)
			for i+1 < len(text) && isPunctuationChar(text[i+1]) && text[i+1] != '\'' {
				punc += string(text[i+1])
				i++
			}
			tokens = append(tokens, punc)
			continue
		}

		current += string(ch)
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
			result = handleCommand(result, tokens[i])
		} else {
			result = append(result, tokens[i])
		}
	}

	return result
}

func isCommand(s string) bool {
	if len(s) < 3 {
		return false
	}
	if s[0] != '(' || s[len(s)-1] != ')' {
		return false
	}
	return true
}

func handleCommand(tokens []string, cmd string) []string {
	inside := cmd[1 : len(cmd)-1]
	action := inside
	count := 1

	commaIndex := -1
	for i := 0; i < len(inside); i++ {
		if inside[i] == ',' {
			commaIndex = i
			break
		}
	}

	if commaIndex != -1 {
		action = strings.TrimSpace(inside[:commaIndex])
		numberPart := strings.TrimSpace(inside[commaIndex+1:])
		n := stringToInt(numberPart)
		if n > 0 {
			count = n
		}
	}

	if action == "hex" {
		idx := findPreviousWord(tokens, len(tokens)-1)
		if idx != -1 {
			value, ok := hexToInt(tokens[idx])
			if ok {
				tokens[idx] = intToString(value)
			}
		}
	}

	if action == "bin" {
		idx := findPreviousWord(tokens, len(tokens)-1)
		if idx != -1 {
			value, ok := binToInt(tokens[idx])
			if ok {
				tokens[idx] = intToString(value)
			}
		}
	}

	if action == "up" || action == "low" || action == "cap" {
		applied := 0
		for i := len(tokens) - 1; i >= 0 && applied < count; i-- {
			if isWord(tokens[i]) {
				if action == "up" {
					tokens[i] = strings.ToUpper(tokens[i])
				}
				if action == "low" {
					tokens[i] = strings.ToLower(tokens[i])
				}
				if action == "cap" {
					tokens[i] = capitalize(tokens[i])
				}
				applied++
			}
		}
	}

	return tokens
}

func findPreviousWord(tokens []string, start int) int {
	for i := start; i >= 0; i-- {
		if isWord(tokens[i]) {
			return i
		}
	}
	return -1
}

func findNextWord(tokens []string, start int) int {
	for i := start; i < len(tokens); i++ {
		if isWord(tokens[i]) {
			return i
		}
	}
	return -1
}

func isWord(s string) bool {
	if s == "" {
		return false
	}
	if s == "'" {
		return false
	}
	if isCommand(s) {
		return false
	}
	if isPunctuation(s) {
		return false
	}
	return true
}

func isPunctuation(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if !isPunctuationChar(s[i]) || s[i] == '\'' {
			return false
		}
	}
	return true
}

func isPunctuationChar(ch byte) bool {
	return ch == '.' || ch == ',' || ch == '!' || ch == '?' || ch == ';' || ch == ':' || ch == '\''
}

func capitalize(s string) string {
	if s == "" {
		return s
	}

	s = strings.ToLower(s)

	first := s[0]
	if first >= 'a' && first <= 'z' {
		first = first - 32
	}

	return string(first) + s[1:]
}

func fixArticles(tokens []string) []string {
	for i := 0; i < len(tokens)-1; i++ {
		if tokens[i] == "a" || tokens[i] == "A" || tokens[i] == "an" || tokens[i] == "An" {
			next := findNextWord(tokens, i+1)
			if next == -1 {
				continue
			}

			needAn := startsWithVowelOrH(tokens[next])

			switch tokens[i] {
			case "a":
				if needAn {
					tokens[i] = "an"
				}
			case "A":
				if needAn {
					tokens[i] = "An"
				}
			case "an":
				if !needAn {
					tokens[i] = "a"
				}
			case "An":
				if !needAn {
					tokens[i] = "A"
				}
			}
		}
	}
	return tokens
}

func startsWithVowelOrH(word string) bool {
	if word == "" {
		return false
	}

	first := word[0]
	if first >= 'A' && first <= 'Z' {
		first = first + 32
	}

	return first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' || first == 'h'
}

func rebuildText(tokens []string) string {
	result := ""
	inQuote := false

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		if tok == "'" {
			if !inQuote {
				if i > 0 {
					prev := tokens[i-1]
					if prev != "'" && !isPunctuation(prev) {
						result += " "
					} else if isPunctuation(prev) {
						result += " "
					}
				}
				result += "'"
				inQuote = true
			} else {
				result += "'"
				inQuote = false
			}
			continue
		}

		if isPunctuation(tok) {
			result = trimTrailingSpace(result)
			result += tok
			continue
		}

		if i > 0 {
			prev := tokens[i-1]
			if prev != "'" {
				result += " "
			}
		}

		result += tok
	}

	return result
}

func trimTrailingSpace(s string) string {
	for len(s) > 0 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s
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

func intToString(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""
	for n > 0 {
		digit := n % 10
		result = string(byte(digit)+'0') + result
		n = n / 10
	}
	return result
}

func binToInt(s string) (int, bool) {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '0' && s[i] != '1' {
			return 0, false
		}
		n = n*2 + int(s[i]-'0')
	}
	return n, true
}

func hexToInt(s string) (int, bool) {
	s = strings.ToLower(s)
	n := 0

	for i := 0; i < len(s); i++ {
		value := -1

		if s[i] >= '0' && s[i] <= '9' {
			value = int(s[i] - '0')
		} else if s[i] >= 'a' && s[i] <= 'f' {
			value = int(s[i]-'a') + 10
		}

		if value == -1 {
			return 0, false
		}

		n = n*16 + value
	}

	return n, true
}
