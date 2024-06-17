package piscine

func basecheck(base string) bool {
	runes := []rune(base)
	len := len(runes)
	if len < 2 {
		return false
	}
	for i := 0; i < len; i++ {
		if runes[i] == '+' || runes[i] == '-' {
			return false
		}
		for j := i + 1; j < len; j++ { 
			if runes[i] == runes[j] {
				return false
			}
		}
	}
	return true
}

func	AtoiBase(s string, base string) int {
	if basecheck(base) == false {
		return 0;
	}
	str := []rune(s)
	runes := []rune(base)
	var j int;
	ans := 0
	for i := 0; i < len(s); i++ {
		for j = 0; j < len(runes); j ++ {
			if str[i] == runes[j] {
				break ;
			}
		}
		ans = ans * len(runes) + j
	}
	return ans
}