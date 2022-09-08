package main

import (
	"fmt"
	"strings"
)

func Inversion(str string){
	newstr := strings.Split(str, " ") 
	j:= len(newstr)-1
	for i:=0; i<len(newstr)/2; i++{
		newstr[i], newstr[j] = newstr[j], newstr[i]
		j--
	}
	fmt.Println(strings.Join(newstr, " "))
}


func main(){
	Inversion("snow dog sun")
}