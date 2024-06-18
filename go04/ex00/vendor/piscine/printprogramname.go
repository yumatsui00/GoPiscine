package piscine

import (
	"os"
	"ft"
)

func PrintProgramName() {
	name := []rune(os.Args[0])
	for i := range name {
		ft.PrintRune(name[i])
	}
}