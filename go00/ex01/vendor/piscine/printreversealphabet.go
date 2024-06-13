package piscine

import "ft"

func PrintReverseaAlphabet(){
	for c := 'z'; c >= 'a'; c--{
		ft.PrintRune(c);
	}
	ft.PrintRune('\n');
}