func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    if len(s) == 0 {
		return 0
	}

	if !(s[0] >= '0' && s[0] <= '9') && s[0] != '-' && s[0] != '+' {
		return 0
	}

	var num int
	sign := 1
	i := 0
	if s[0] == '-' {
		i = 1
		sign = -1
	}

	if s[0] == '+' {
		i = 1
	}

	for ; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			break
		}

		num = (num * 10) + (int(s[i]) - 48)
		if sign == -1 && num*sign <= math.MinInt32 {
			return math.MinInt32
		}

		if sign == 1 && num >= math.MaxInt32 {
			return math.MaxInt32
		}

	}
	return num * sign
}