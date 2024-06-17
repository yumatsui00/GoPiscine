package main

import (
	"piscine"
	"fmt"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("あalut!", "あ"))
	fmt.Println(piscine.Index("Ola!", "h01"))
}