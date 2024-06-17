package piscine

func	Join(strs []string, sep string) string {
	size := len(strs)
	var	ret	string
	for i := 0; i < size; i++ {
		ret += strs[i]
		if i < size - 1 {
			ret += sep
		}
	}
	return ret
}