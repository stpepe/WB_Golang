package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"errors"
)

func unpack(str string) (string, error) {
	arrWords := strings.Split(str, "")
	for i, _ := range arrWords{
		if i != len(arrWords)-1{
			_, err1 := strconv.Atoi(arrWords[i])
			_, err2 := strconv.Atoi(arrWords[i+1])
			if err1 == nil && err2 == nil{
				return "", errors.New("Некорректная строка!")
			}
		}
	}
	var sb strings.Builder
	for i, el := range arrWords{
		if intVar, err := strconv.Atoi(el); err == nil{
			if i == 0{
				return "", errors.New("Некорректная строка!")
			} 
			for a:=1; a < intVar; a++{
				sb.WriteString(arrWords[i-1])
			}
		} else {
			sb.WriteString(el)
		}
	}
	return sb.String(), nil
}

func main() {
	res, err := unpack("adf22s2")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(res)
}
