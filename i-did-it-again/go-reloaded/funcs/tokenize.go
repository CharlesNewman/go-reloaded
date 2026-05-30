package funcs

import "strings"

func Tokenize(text string) []string {

	tokens := strings.Fields(text)

	return tokens
}
