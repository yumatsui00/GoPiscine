package piscine

import "ft"

func nv( void ) bool {
	ft.PrintRune('N')
	ft.PrintRune('V')
	return false
}

func basecheck(base string) bool {
	runes := []rune(base)
	len := len(runes)
	if len < 2 {
		return nv()
	}
	for i := 0; i < len; i++ {
		if (runes[i] == '+' || runes[i] == '-')
			return nv()
		j := i
		while (++j < len)
		{
			if (runes[i] == runes[j])
				return nv()
		}
	}
	return true
}

func	PrintNbrBase(nbr int, base string) {
	if (basecheck(base) == false)
		return
	runes = []rune(base)
	if nb < 0 {
		ft.PrintRune('-')
	} else if ()
}