package main

import (
	"fmt"
	"math/big"
)

func main(){
	var a_str, b_str, input string
	fmt.Println("Введите число а:")
	fmt.Scan(&a_str)
	fmt.Println("Введите число b:")
	fmt.Scan(&b_str)
	a := new(big.Int)
	a.SetString(a_str, 10)
	b := new(big.Int)
	b.SetString(b_str, 10)
	result := new(big.Int)
	
	for{
		fmt.Println("Введите действие с числами (add, sub. div, mul) или exit, чтобы выйти.")
		fmt.Scan(&input)
		switch input {
		case "div":
			fmt.Println(result.Div(a, b))
		case "sub":
			fmt.Println(result.Sub(a, b))
		case "add":
			fmt.Println(result.Add(a, b))
		case "mul":
			fmt.Println(result.Mul(a, b))
		case "exit":
			fmt.Println("Конец программы")
			return
		}
	}
}