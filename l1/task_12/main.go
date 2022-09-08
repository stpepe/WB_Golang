package main 

import (
	"fmt"
)

func CreateSet(mass *[]string) []string {
	tmp := map[string]int{}
	res := []string{}
	for _, el := range *mass{
		if _, ok := tmp[el]; !ok{
			tmp[el] = 0
			res = append(res, el)
		}
	}
	return res
}

func main(){
	mass := []string{"cat", "cat", "dog", "cat", "tree"}

	res := CreateSet(&mass)

	fmt.Println(res)
}