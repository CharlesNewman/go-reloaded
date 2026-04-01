package funcs

func Tokenize(text string) []string {
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
