package funcs

func BuildText(tokens []string) string {
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
				result = removeLastSpace(result)
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

			if inQuote {
				if last != '\'' && last != '\n' && last != ' ' {
					result += " "
				}
			} else {
				if last != '\n' && last != ' ' && last != '\'' {
					result += " "
				}
				if last == '\'' {
					result += " "
				}
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
