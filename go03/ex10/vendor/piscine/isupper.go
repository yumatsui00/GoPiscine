package piscine

func	isupper(a rune) bool {
	if a < 'A' || a > 'Z' {
		return (false)
	}
	return (true)
}

func	IsUpper(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if isupper(runes[i]) == false {
			return false
		} 
	}
	return true
}