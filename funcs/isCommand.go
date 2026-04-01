package funcs

func isCommand(s string) bool {
	if len(s) < 2 {
		return false
	}
	return s[0] == '(' && s[len(s)-1] == ')'
}
