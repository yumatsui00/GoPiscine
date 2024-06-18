package piscine

import (
	"os"
	"ft"
)

func	PrintParams() {
	for i := 1; i < len(os.Args); i++ {
		name := []rune(os.Args[i])
		for j := range name {
			ft.PrintRune(name[j])
		}
		ft.PrintRune('\n')
	}
}