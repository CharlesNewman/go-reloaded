package funcs

import (
	"strconv"
	"strings"
)

func applyUp(word string) string {
	return strings.ToUpper(word)
}

func applyLow(word string) string {
	return strings.ToLower(word)
}

func applyCap(word string) string {

	buildCap := ""
	finalWord := ""

	if word == "" {
		return word
	}

	buildWord := word[1:]

	if word[0] >= 'a' && word[0] <= 'z' {
		buildCap = string(word[0] - 32)
		finalWord = buildCap + strings.ToLower(buildWord)
	} else if word[0] >= 'A' && word[0] <= 'Z' {
		finalWord = string(word[0]) + strings.ToLower(buildWord)
	}
	return finalWord
}

func applyBin(word string) string {

	num, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(num, 10)
}

func applyHex(word string) string {

	num, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(num, 10)
}

func SelectCommand(word string, command string) string {

	newWord := word

	if command == "(hex)" {
		newWord = applyHex(word)
	}
	if command == "(bin)" {
		newWord = applyBin(word)
	}
	if command == "(up)" || command == "(up," {
		newWord = applyUp(word)
	}
	if command == "(low)" || command == "(low," {
		newWord = applyLow(word)
	}
	if command == "(cap)" || command == "(cap," {
		newWord = applyCap(word)
	}
	return newWord
}

func FixPunctuation(line string) string {
	// Step 1: remove spaces BEFORE punctuation
	for i := 0; i < len(line); i++ {
		ch := line[i]
		if IsPunctuation(ch) {
			line = strings.ReplaceAll(line, " "+string(ch), string(ch))
		}
	}

	// Step 2: add space AFTER punctuation when needed
	build := ""

	for i := 0; i < len(line); i++ {
		ch := line[i]
		build += string(ch)

		next := i + 1

		if IsPunctuation(ch) &&
			next < len(line) &&
			line[next] != ' ' &&
			!IsPunctuation(line[next]) {
			build += " "
		}
	}

	return build
}

func IsPunctuation(ch byte) bool {
	return ch == '.' || ch == ',' || ch == '!' || ch == '?' || ch == ':' || ch == ';'
}

func FixQuotes(line string) string {
	result := ""
	sentence := ""
	insideQuote := false

	for i := 0; i < len(line); i++ {
		if line[i] == '\'' {
			if insideQuote {
				// closing quote
				sentence = strings.TrimSpace(sentence)
				result += sentence + "'"
				sentence = ""
				insideQuote = false
			} else {
				// opening quote
				result += "'"
				insideQuote = true
			}
		} else {
			if insideQuote {
				sentence += string(line[i])
			} else {
				result += string(line[i])
			}
		}
	}

	return result
}

func FixArticle(line string) string {
	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		if i+1 >= len(words) {
			continue
		}

		nextWord := words[i+1]
		if nextWord == "" {
			continue
		}

		first := strings.ToLower(string(nextWord[0]))

		needsAn := first == "a" || first == "e" || first == "i" || first == "o" || first == "u" || first == "h"

		if words[i] == "a" && needsAn {
			words[i] = "an"
		} else if words[i] == "A" && needsAn {
			words[i] = "An"
		} else if words[i] == "an" && !needsAn {
			words[i] = "a"
		} else if words[i] == "An" && !needsAn {
			words[i] = "A"
		}
	}

	return strings.Join(words, " ")
}
