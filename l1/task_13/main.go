package main 

import (
	"fmt"
)

func main(){
	var first int = 12
	var second int = 33

	fmt.Printf("First = %d; Second = %d;\n", first, second)

	first, second = second, first
	
	fmt.Printf("First = %d; Second = %d;\n", first, second)
}