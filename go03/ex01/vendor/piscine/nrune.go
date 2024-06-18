package piscine

import "ft"

func	NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	runes := []rune(s)
	if n > ft.StrLen(s) {
		return 0
	}
	return runes[n - 1]
}