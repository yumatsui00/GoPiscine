package ft

import (
	"errors"
	"os"
	"unicode/utf8"
)

// PrintRune prints a single rune (Unicode code point) and returns any error
// if the encoding or the writing fails.
func PrintRune(r rune) error {
	l := utf8.RuneLen(r)
	if l == -1 {
		return errors.New("The rune is not a valid value to encode in UTF-8")
	}
	p := make([]byte, l)
	utf8.EncodeRune(p, r)
	_, err := os.Stdout.Write(p)
	return err
}

func PrintInt(i int){
	if i < 0 {
		PrintRune('-')
		PrintInt(-1 * i)
	}
	else if i >= 10 {
		PrintInt(i / 10)
		PrintRune(rune(i % 10) + '0')
	}
	else {
		PrintRune(rune(i % 10) + '0')
	}

}

func Ft_Atoi(str string){
	res := 0
	flag := 1

	for i := 0; str[i] && (str[i] == ' ' || str[i] == '+') {
		i ++;
	}
	if str[i] == '-' {
		flag = -1;
	}
	for i = 0; str[i] && str[i] >= '0' && str[i] <= '9' {
		res += res * 10 + int(str[i] - '0')
	}
	return res * flag
}