package piscine

import "ft"

func	ToLower(s string) string {
	runes := []rune(s)
	for i := 0; i < ft.StrLen(s); i++ {
		if (runes[i] >= 'A' && runes[i] <= 'Z') {
			runes[i] += 'a' - 'A';
		}
	}
	ret := string(runes)
	return ret
}