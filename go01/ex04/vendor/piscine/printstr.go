package piscine

import "ft"

func PrintStr(a string){
	for i := 0; i < len(a); i++{
		ft.PrintRune(rune(a[i]))
	}
}