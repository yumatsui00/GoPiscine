package piscine

func	islower(a rune) bool {
	if a < 'a' || a > 'z' {
		return (false)
	}
	return (true)
}

func	IsLower(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if islower(runes[i]) == false {
			return false
		} 
	}
	return true
}