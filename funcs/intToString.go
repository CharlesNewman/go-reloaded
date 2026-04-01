package funcs

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
