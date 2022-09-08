package main 

import (
	"fmt"
)

func MakeSet(first, second *[]int) []int {
	tmp := map[int]int{}
	res := []int{}

	for _, el := range *first{
		if _, ok := tmp[el]; !ok{
			tmp[el] = 0
		}
	}

	for _, el := range *second{
		if _, ok := tmp[el]; ok{
			res = append(res, el)
		}
	}
	return res
}

func main(){
	first := []int{1, 2, 5, 5, 6, 3, 12}
	second := []int{5, 3, 7, 8, 12}
	
	res := MakeSet(&first, &second)

	fmt.Println(res)
}