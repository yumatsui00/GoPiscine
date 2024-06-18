package piscine

func	islower(a rune) bool {
	if a < 'a' || a > 'z' {
		return (false)
	}
	return (true)
}

func	isupper(a rune) bool {
	if a < 'A' || a > 'Z' {
		return (false)
	}
	return (true)
}

func	isdigit(a rune) bool {
	if a < '0' || a > '9' {
		return (false)
	}
	return (true)
}

func	IsNumeric(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if isdigit(runes[i]) == false {
			return false
		} 
	}
	return true
}