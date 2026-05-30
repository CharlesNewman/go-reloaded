package funcs

import (
	"strconv"
	"strings"
)

func ProcessText(text string) string {

	line := strings.Split(text, "\n")

	for i := 0; i < len(line); i++ {
		line[i] = ProcessLine(line[i])
	}
	return strings.Join(line, "\n")
}

func ProcessLine(line string) string {

	tokens := Tokenize(line)
	command := ""
	fixedText := ""

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(hex)" || tokens[i] == "(bin)" || tokens[i] == "(up)" || tokens[i] == "(low)" || tokens[i] == "(cap)" {
			command = tokens[i]
			if i > 0 {
				tokens[i-1] = SelectCommand(tokens[i-1], command)
				tokens[i] = ""
			}
		}
	}

	for x := 0; x < len(tokens); x++ {
		y := x + 1
		if tokens[x] == "(up," || tokens[x] == "(low," || tokens[x] == "(cap," {
			if y < len(tokens) && strings.HasSuffix(tokens[y], ")") {
				numberText := strings.TrimSuffix(tokens[y], ")")
				num, err := strconv.Atoi(numberText)
				if err != nil {
					return line
				}
				start := x - num
				if start < 0 {
					start = 0
				}
				command = tokens[x]
				tokens[x] = ""
				tokens[y] = ""
				for j := start; j < x; j++ {
					tokens[j] = SelectCommand(tokens[j], command)
				}

			}
		}

	}
	cleanTokens := []string{}
	for j := 0; j < len(tokens); j++ {
		if tokens[j] != "" {
			cleanTokens = append(cleanTokens, tokens[j])
		}
	}
	fixedText = strings.Join(cleanTokens, " ")
	fixedText = FixPunctuation(fixedText)
	fixedText = FixQuotes(fixedText)
	fixedText = FixArticle(fixedText)

	return fixedText
}
