package main

import (
	"fmt"
	"sync"
)

func Factory(first, second chan int){
	for msg := range first{
		second <- msg*2
	}
	close(second)
}

func Printer(second chan int){
	for msg:= range second{
		fmt.Println(msg)
	}
}

func main(){
	mass := [5]int{1, 2, 3, 4, 5} 
	first := make(chan int)
	second := make(chan int)
	var wg sync.WaitGroup

	go func(){
		wg.Add(1)
		Factory(first, second)
		wg.Done()
	}()

	go func(){
		wg.Add(1)
		Printer(second)
		wg.Done()
	}()

	for _, el := range mass{
		first <- el
	}
	close(first)

	wg.Wait()
}