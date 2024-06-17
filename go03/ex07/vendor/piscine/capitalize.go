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

func	Capitalize(s string) string {
	upperflag := true
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if upperflag == true && islower(runes[i]) == true {
			runes[i] += 'A' - 'a'
		}
		upperflag = false
		if (isupper(runes[i]) == false && islower(runes[i]) == false && isdigit(runes[i]) == false) {
			upperflag = true
		}
	} 
	ret := string(runes)
	return ret
}