package piscine

func	BasicJoin(elems []string) string {
	var ret	string;
	for i := range elems {
		ret += elems[i]
	}
	return ret
}