package main

import (
	"piscine"
	"fmt"
)

func main(){
	a := 0
	b := 1
	piscine.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}