package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
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
		fmt.Println("Error reading input file:", err)
		return
	}

	result := processText(string(data))

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
	// This regex finds:
	// - commands like (up), (low, 2), (hex), (bin)
	// - words/numbers
	// - punctuation groups like ..., !!, !?, etc
	// - single quote '
	re := regexp.MustCompile(`\((?:hex|bin|up|low|cap)(?:,\s*\d+)?\)|[A-Za-z0-9]+(?:'[A-Za-z0-9]+)?|\.{3}|[.,!?;:]+|'`)
	return re.FindAllString(text, -1)
}

func applyCommands(tokens []string) []string {
	var result []string

	for _, tok := range tokens {
		if isCommand(tok) {
			result = handleCommand(result, tok)
		} else {
			result = append(result, tok)
		}
	}

	return result
}

func isCommand(s string) bool {
	return strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")")
}

func handleCommand(tokens []string, cmd string) []string {
	re := regexp.MustCompile(`^\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)$`)
	matches := re.FindStringSubmatch(cmd)
	if matches == nil {
		return tokens
	}

	action := matches[1]
	count := 1

	if matches[2] != "" {
		n, err := strconv.Atoi(matches[2])
		if err == nil && n > 0 {
			count = n
		}
	}

	switch action {
	case "hex":
		idx := findPreviousWord(tokens, len(tokens)-1)
		if idx != -1 {
			if value, err := strconv.ParseInt(tokens[idx], 16, 64); err == nil {
				tokens[idx] = strconv.FormatInt(value, 10)
			}
		}
	case "bin":
		idx := findPreviousWord(tokens, len(tokens)-1)
		if idx != -1 {
			if value, err := strconv.ParseInt(tokens[idx], 2, 64); err == nil {
				tokens[idx] = strconv.FormatInt(value, 10)
			}
		}
	case "up", "low", "cap":
		applied := 0
		for i := len(tokens) - 1; i >= 0 && applied < count; i-- {
			if isWord(tokens[i]) {
				switch action {
				case "up":
					tokens[i] = strings.ToUpper(tokens[i])
				case "low":
					tokens[i] = strings.ToLower(tokens[i])
				case "cap":
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
	for _, r := range s {
		if !strings.ContainsRune(".,!?;:", r) {
			return false
		}
	}
	return true
}

func capitalize(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func fixArticles(tokens []string) []string {
	for i := 0; i < len(tokens)-1; i++ {
		if strings.EqualFold(tokens[i], "a") {
			next := findNextWord(tokens, i+1)
			if next != -1 && startsWithVowelOrH(tokens[next]) {
				if tokens[i] == "A" {
					tokens[i] = "An"
				} else {
					tokens[i] = "an"
				}
			}
		}
	}
	return tokens
}

func findNextWord(tokens []string, start int) int {
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

	r := unicode.ToLower([]rune(word)[0])
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h'
}

func rebuildText(tokens []string) string {
	var b strings.Builder
	inQuote := false

	for i, tok := range tokens {
		if tok == "'" {
			if !inQuote {
				// opening quote
				if i > 0 {
					prev := tokens[i-1]
					if prev != "'" && !isPunctuation(prev) {
						b.WriteString(" ")
					} else if isPunctuation(prev) {
						b.WriteString(" ")
					}
				}
				b.WriteString("'")
				inQuote = true
			} else {
				// closing quote
				b.WriteString("'")
				inQuote = false
			}
			continue
		}

		if isPunctuation(tok) {
			trimTrailingSpace(&b)
			b.WriteString(tok)
			continue
		}

		if i > 0 {
			prev := tokens[i-1]
			if prev != "'" {
				b.WriteString(" ")
			}
		}

		b.WriteString(tok)
	}

	return b.String()
}

func trimTrailingSpace(b *strings.Builder) {
	s := b.String()
	if strings.HasSuffix(s, " ") {
		newStr := strings.TrimRight(s, " ")
		b.Reset()
		b.WriteString(newStr)
	}
}
