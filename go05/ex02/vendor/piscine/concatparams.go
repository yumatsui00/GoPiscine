package piscine

func	ConcatParams(args []string) string {
	count := 0
	for range args {
		count++
	}
	ret := ""
	first := true
	for i := 0; i < count; i++ {
		if first == true {
			first = false
		} else {
			ret += "\n"
		}
		ret += args[i]
	}
	return ret
}