package main

import (
	"fmt"
	"time"
	"math/rand"
)

func Sender(ch chan int){
	for{
		ch <- rand.Intn(1000)
		time.Sleep(time.Second*1)
	}
}

func Reader(ch chan int, mass *[]int){
	for{
		tmp := <- ch
		*mass = append(*mass, tmp)
	}
}

func main(){
	ch := make(chan int)
	mass := []int{}
	go Sender(ch)
	go Reader(ch, &mass)

	time.Sleep(time.Second*5)
	close(ch)
	fmt.Println(mass)	
}