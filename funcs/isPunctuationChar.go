package funcs

func isPunctuationChar(ch byte) bool {
	if ch == '.' || ch == ',' || ch == '!' || ch == '?' || ch == ':' || ch == ';' {
		return true
	}
	return false
}
