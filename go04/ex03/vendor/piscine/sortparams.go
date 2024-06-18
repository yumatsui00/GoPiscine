package piscine

import (
	"os"
	"ft"
)

func SortParams() {
	args := os.Args[1:]
	argc := 0
	for range args {
		argc++
	}
    for i := 0; i < argc - 1; i++ {
        for j := 0; j < argc - i - 1; j++ {
            if args[j] > args[j + 1] {
                tmp := args[j]
                args[j] = args[j + 1]
                args[j + 1] = tmp
            }
        }
    }
	for i := 0;  i < argc; i++ {
		name := []rune(args[i])
		for j := range name {
			ft.PrintRune(name[j])
		}
		ft.PrintRune('\n')
	}
}
