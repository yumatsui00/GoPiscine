package piscine

import (
	"os"
	"ft"
)

func	PrintParams() {
	first := true
	for i := range os.Args  {
		if first {
			first = false
		} else {
			name := []rune(os.Args[i])
			for j := range name {
				ft.PrintRune(name[j])
			}
		ft.PrintRune('\n')
		}
	}
}