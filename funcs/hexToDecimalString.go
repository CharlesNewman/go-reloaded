package funcs

func hexToDecimalString(s string) string {
	value := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]
		digit := 0

		if ch >= '0' && ch <= '9' {
			digit = int(ch - '0')
		} else if ch >= 'A' && ch <= 'F' {
			digit = int(ch-'A') + 10
		} else if ch >= 'a' && ch <= 'f' {
			digit = int(ch-'a') + 10
		} else {
			return s
		}

		value = value*16 + digit
	}

	return intToString(value)
}
