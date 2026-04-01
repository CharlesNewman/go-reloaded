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

func intToString(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""

	for n > 0 {
		digit := n % 10
		result = string(byte(digit+'0')) + result
		n = n / 10
	}

	return result
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
