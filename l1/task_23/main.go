package main

import (
	"fmt"
)

func DeleteElement(mass []int, i int) []int {
	return append(mass[:i], mass[i+1:]...)
}

func main(){
	mass := []int{1, 2, 3, 4, 5}
	mass = DeleteElement(mass, 2)
	fmt.Println(mass)
}