package piscine

import "ft"

func nv() bool {
	ft.PrintRune('N')
	ft.PrintRune('V')
	return false
}

func basecheck(base string) bool {
	runes := []rune(base)
	len := ft.StrLen(base)
	if len < 2 {
		return nv()
	}
	for i := 0; i < len; i++ {
		if runes[i] == '+' || runes[i] == '-' {
			return nv()
		}
		for j := i + 1; j < len; j++ { 
			if runes[i] == runes[j] {
				return nv()
			}
		}
	}
	return true
}

func	PrintNbrBase(nbr int, base string) {
	if basecheck(base) == false {
		return
	}
	runes := []rune(base)
	if nbr < 0 {
		ft.PrintRune('-')
		PrintNbrBase(-1 * nbr, base)
	} else if nbr >= ft.StrLen(base) {
		PrintNbrBase(nbr / ft.StrLen(base), base)
		ft.PrintRune(runes[nbr % ft.StrLen(base)])
	} else {
		ft.PrintRune(runes[nbr])
	}
}