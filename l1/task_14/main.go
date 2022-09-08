package main

import (
	"fmt"
)

func TheType(something interface{}) {
	fmt.Printf("Type is %T\n", something)
}

func main() {
	intt := 7
	stringg := "golang"
	booll := true
	chani := make(chan int)
	chans := make(chan string)
	obj := struct{}{}

	TheType(intt)
	TheType(stringg)
	TheType(booll)
	TheType(chani)
	TheType(chans)
	TheType(obj)
}