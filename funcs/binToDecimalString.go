package funcs

func binToDecimalString(s string) string {
	value := 0

	for i := 0; i < len(s); i++ {
		if s[i] != '0' && s[i] != '1' {
			return s
		}
		value = value*2 + int(s[i]-'0')
	}

	return intToString(value)
}
