package main

import (
	"piscine"
	"fmt"
)

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.BasicJoin(elems))
}