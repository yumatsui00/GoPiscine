package piscine

import "ft"

func IsNegative(n int){
	if n >= 0 {
		ft.PrintRune('F')
	} else {
		ft.PrintRune('T')
	}
	ft.PrintRune('\n');
}