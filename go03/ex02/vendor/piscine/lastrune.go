package piscine

import "ft"

func	LastRune(s string) rune {
	runes := []rune(s)
	if ft.StrLen(s) <= 0 {
		return 0
	}
	return runes[ft.StrLen(s) - 1]
}