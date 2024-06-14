package piscine

import "ft"

func checkqueen(ans [8]int, row int) bool {
	for i := 0; i < row; i++ {
		if ans[i] == ans[row] {
			return false
		} else if ans[i] - ans[row] == row - i || ans[i] - ans[row] == i - row {
			return false
		}
	}
	return true
}

func rec(ans [8]int, row int) {
	if (row == 8) {
		for i := 0; i < 8; i++ {
			ft.PrintRune(rune(ans[i] + '1'))
		}
		ft.PrintRune('\n')
		return ;
	}
	for column := 0; column < 8; column++ {
		ans[row] = column
		if checkqueen(ans, row) == true {
			rec(ans, row + 1)
		}
	}
}

func EightQueens() {
	var ans [8]int
	rec(ans, 0)
}