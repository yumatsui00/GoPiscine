package piscine

func	BasicJoin(elems []string) string {
	size := len(elems)
	var ret	string;
	for i := 0; i < size; i++ {
		ret += elems[i]
	}
	return ret
}