package main

import (
	"fmt"
	"sync"
)

func Squaring(num int, sum *int, mu *sync.Mutex, wt *sync.WaitGroup){
	mu.Lock()
	defer mu.Unlock()
	defer wt.Done()
	*sum += num*num 
	return
}

func main(){
	var mu sync.Mutex
	var wt sync.WaitGroup
	mass := []int{2, 4, 6, 8, 10}
	wt.Add(len(mass))
	var sum int
	for _, el := range mass{
		go Squaring(el, &sum, &mu, &wt)
	}
	wt.Wait()
	fmt.Printf("Sum of squares: %d\n",sum)
}