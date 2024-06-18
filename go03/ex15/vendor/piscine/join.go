package piscine

func	Join(strs []string, sep string) string { 
	var	ret	string
	first := true
	for i := range strs {
		if first == true {
			ret += strs[i]
			first = false 
		} else {
			ret += sep
			ret += strs[i]
		}
	}
	return ret
}