package piscine

func	isprint(a rune) bool {
	if a < 32 || a > 126 {
		return true
	}
	return false
}

//! まだだ
func	IsPrintable(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if isprint(runes[i]) == false {
			return false
		} 
	}
	return true
}