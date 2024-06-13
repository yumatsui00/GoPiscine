package piscine

import "ft"

func printnum(i int) {
	if i < 10 {
		ft.PrintRune('0')
		ft.PrintRune(rune(i + '0'))
	} else {
		ft.PrintRune(rune(i % 10 + '0'))
		ft.PrintRune(rune(i / 10 + '0'))
	}
}

func rec(i int) {
	if i == 99 {
		return
	}
	for j := i + 1; j < 100; j++ {
		printnum(i)
		ft.PrintRune(' ')
		printnum(j)
		if (i != 98) {
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
	}
	rec(i + 1)
}

func PrintComb2(){
	rec(0)
}