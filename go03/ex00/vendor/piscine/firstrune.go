package piscine

import "ft"

func FirstRune(s string) rune {
	runes := []rune(s)
	if ft.StrLen(s) <= 0 {
		return 0
	}
	return runes[0]
}