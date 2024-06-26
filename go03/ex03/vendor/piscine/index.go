package piscine

import "ft"

func	Index(s string, toFind string) int {
	sRune := []rune(s)
	toFindRune := []rune(toFind)
	sLen := ft.StrLen(s)
	toFindLen := ft.StrLen(toFind)
	if sLen <= 0 || toFindLen <= 0 {
		return -1
	}
	for i := 0; i < sLen; i++ {
		if sRune[i] == toFindRune[0] {
			for j := 0; sRune[i + j] == toFindRune[j]; j++ {
				if j == toFindLen - 1 {
					return i
				}
			}
		}
	}
	return -1
}