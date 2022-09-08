package main

import (
	"fmt"
)

func Inversion(str string){
	newstr := []rune(str)
	j:= len(newstr)-1
	for i:=0; i<len(newstr)/2; i++{
		newstr[i], newstr[j] = newstr[j], newstr[i]
		j--
	}
	fmt.Println(string(newstr))
}

func main(){
	Inversion("главрыба")
}