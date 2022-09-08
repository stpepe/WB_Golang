package main

import (
	"fmt"
)

func createHugeString(size int) string {
	res := []byte{}
	for i:=0; i<size; i++{
		res = append(res, ' ')
	}
	return string(res)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100]
}

func main() {
	var justString string = someFunc()
	fmt.Println(justString)
	fmt.Println(len(justString))
}

