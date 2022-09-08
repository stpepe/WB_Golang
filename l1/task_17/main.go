package main

import (
	"fmt"
)

func Printer(res *int){
	fmt.Printf("Число %d найдено!\n", *res)
}

func BinarySearch(mass []int, num *int){
	if len(mass) == 1{
		Printer(&mass[0])
		return 
	}

	half := len(mass)/2

	if mass[half] > *num {
		BinarySearch(mass[:half], num)
	} else if mass[half] == *num {
		Printer(&mass[half])
		return 
	} else {
		BinarySearch(mass[half+1:], num)
	}

	return 
}

func main(){
	mass := []int{1, 2, 5, 7, 17, 22, 35, 48, 54}
	num := 5
	BinarySearch(mass, &num)
	
	// for _, v := range mass{
	// 	BinarySearch(mass, &v)
	// }
}