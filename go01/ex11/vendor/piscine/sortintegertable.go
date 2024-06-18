package piscine

func SortIntegerTable(table []int){
	len := ft.StrLen(table)
	for i := 0; i < len - 1; i++ {
		for j := 0; j < len - 1; j++ {
			if table[j] > table[j + 1] {
				tmp := table[j]
				table[j] = table[j + 1]
				table[j + 1] = tmp
			}
		}
	}
}