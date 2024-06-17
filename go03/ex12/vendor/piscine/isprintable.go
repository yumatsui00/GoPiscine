package piscine

func	isprint(a rune) bool {
	if a < 32 && a >= 0 {
		return true
	}
	return false
}

//! まだだ
func	IsPrintable(s string) bool {
	runes := []rune(s)
	for i := range runes{
		if isprint(runes[i]) == true {
			return false
		} 
	}
	return true
}