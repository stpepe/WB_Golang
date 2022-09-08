package main

import (
	"fmt"
	"strings"
)

func UniqueString(mass []string) bool{
	for i, _ := range mass{
		for j, _ := range mass{
			if (mass[i] == mass[j]) && (i != j){
				return false
			}
		}
	}
	return true
}

func main(){
	var str string
	fmt.Println("Введите строку для проверки уникальности:")
	fmt.Scan(&str)
	mass := strings.Split(str, "")
	fmt.Println(UniqueString(mass))
}