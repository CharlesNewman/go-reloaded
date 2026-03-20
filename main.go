package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong input expected format: go run . input.txt output.txt")
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
	text = applyCommandsDirect(text)
	text = fixArticlesDirect(text)
	return strings.TrimSpace(text)
}

func applyCommandsDirect(text string) string {
	for {
		start := -1
		end := -1

		for i := 0; i < len(text); i++ {
			if text[i] == '(' {
				start = i
				break
			}
		}

		if start == -1 {
			break
		}

		for i := start; i < len(text); i++ {
			if text[i] == ')' {
				end = i
				break
			}
		}

		if end == -1 {
			break
		}

		cmd := text[start : end+1]
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

		if action == "hex" || action == "bin" {
			wordStart, wordEnd := findPreviousWordBounds(text, start-1)
			if wordStart != -1 {
				word := text[wordStart : wordEnd+1]

				if action == "hex" {
					value, ok := hexToInt(word)
					if ok {
						text = text[:wordStart] + intToString(value) + text[wordEnd+1:start] + text[end+1:]
					} else {
						text = text[:start] + text[end+1:]
					}
				}

				if action == "bin" {
					value, ok := binToInt(word)
					if ok {
						text = text[:wordStart] + intToString(value) + text[wordEnd+1:start] + text[end+1:]
					} else {
						text = text[:start] + text[end+1:]
					}
				}
			} else {
				text = text[:start] + text[end+1:]
			}
			continue
		}

		if action == "up" || action == "low" || action == "cap" {
			newText := text
			pos := start - 1

			for applied := 0; applied < count; applied++ {
				wordStart, wordEnd := findPreviousWordBounds(newText, pos)
				if wordStart == -1 {
					break
				}

				word := newText[wordStart : wordEnd+1]

				if action == "up" {
					word = strings.ToUpper(word)
				}
				if action == "low" {
					word = strings.ToLower(word)
				}
				if action == "cap" {
					word = capitalize(word)
				}

				newText = newText[:wordStart] + word + newText[wordEnd+1:]
				pos = wordStart - 1
			}

			text = newText[:start] + newText[end+1:]
		} else {
			text = text[:start] + text[end+1:]
		}
	}

	return text
}

func findPreviousWordBounds(text string, pos int) (int, int) {
	for pos >= 0 && (text[pos] == ' ' || text[pos] == '\n' || text[pos] == '\t') {
		pos--
	}

	if pos < 0 {
		return -1, -1
	}

	if isPunctuationChar(text[pos]) {
		for pos >= 0 && isPunctuationChar(text[pos]) {
			pos--
		}
		for pos >= 0 && (text[pos] == ' ' || text[pos] == '\n' || text[pos] == '\t') {
			pos--
		}
	}

	if pos < 0 {
		return -1, -1
	}

	end := pos

	for pos >= 0 &&
		text[pos] != ' ' &&
		text[pos] != '\n' &&
		text[pos] != '\t' &&
		!isPunctuationChar(text[pos]) &&
		text[pos] != '(' &&
		text[pos] != ')' {
		pos--
	}

	start := pos + 1

	if start > end {
		return -1, -1
	}

	return start, end
}

func fixArticlesDirect(text string) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	for i := 0; i < len(words)-1; i++ {
		current := words[i]
		next := cleanWord(words[i+1])

		if next == "" {
			continue
		}

		needAn := startsWithVowelOrH(next)

		switch current {
		case "a":
			if needAn {
				words[i] = "an"
			}
		case "A":
			if needAn {
				words[i] = "An"
			}
		case "an":
			if !needAn {
				words[i] = "a"
			}
		case "An":
			if !needAn {
				words[i] = "A"
			}
		}
	}

	return strings.Join(words, " ")
}

func cleanWord(word string) string {
	start := 0
	end := len(word) - 1

	for start <= end && isPunctuationChar(word[start]) {
		start++
	}
	for end >= start && isPunctuationChar(word[end]) {
		end--
	}

	if start > end {
		return ""
	}

	return word[start : end+1]
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

func fixPunctuationSpacing(text string) string {
	result := ""
	inQuote := false

	for i := 0; i < len(text); i++ {
		ch := text[i]

		if ch == ' ' || ch == '\n' || ch == '\t' {
			if len(result) > 0 && result[len(result)-1] != ' ' {
				result += " "
			}
			continue
		}

		if ch == '\'' {
			if !inQuote {
				if len(result) > 0 && result[len(result)-1] != ' ' {
					result += " "
				}
				result += "'"
				inQuote = true
			} else {
				result = trimTrailingSpace(result)
				result += "'"
				inQuote = false
			}
			continue
		}

		if isPunctuationChar(ch) {
			result = trimTrailingSpace(result)
			result += string(ch)
			continue
		}

		if len(result) > 0 && result[len(result)-1] != ' ' && result[len(result)-1] != '\'' {
			result += " "
		}

		result += string(ch)
	}

	return trimTrailingSpace(result)
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
