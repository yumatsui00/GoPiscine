package piscine

import (
	"os"
	"ft"
)

func RevParams() {
	argc := 0
	for range os.Args  {
		argc++;
	}
	for argc > 1 {
		argc--
		name := []rune(os.Args[argc])
		for j := range name {
			ft.PrintRune(name[j])
		}
		ft.PrintRune('\n')
	}
}

// func	RevParams() {
// 	len := 0
// 	for i := 1; os.Args[i] != ""; i++ {
// 		len ++
// 	}
// 	for i := len - 1; i > 1; i-- {
// 		name := []rune(os.Args[i])
// 		for j := range name {
// 			ft.PrintRune(name[j])
// 		}
// 		ft.PrintRune('\n')
// 	}
// }