package piscine

func	SplitWhiteSpaces(s string) []string {
	count := 0
	for i := range s {
		if i == ' ' {
			count ++
		}
	}
	ret = make([]string, count + 1)

}