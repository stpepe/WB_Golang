package main

import (
	"fmt"
	"strconv"
)

func main(){
	var num int64
	var i int

	fmt.Println("Введите число типа int64: ")
	fmt.Scan(&num)
	bin := strconv.FormatInt(num, 2)
	bin_rune := []rune(bin)
	fmt.Println(num, "- в двоичной системе:", string(bin_rune))
	fmt.Println("Введите номер бита, который нужно инвертировать: ")
	fmt.Scan(&i)
	i = len(bin_rune)-i
	if bin_rune[i] == '0'{
		bin_rune[i] = '1'
	} else {
		bin_rune[i] = '0'
	}
	fmt.Println("После инвертирования в двоичной системе:", string(bin_rune))
	num = 0
	for s, el := range string(bin_rune){
		var degree int64 = 1
			for l:=0; l < (len(bin_rune)-1-s); l++{
				if s == (len(bin_rune)-1){
					return
				}
				degree = degree * 2
			}
		if el == '1'{
			num = num + degree
		}
	}
	fmt.Println("После инвертирования в десятичной системе:", num)

}