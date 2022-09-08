package main

//**********************Вариант 1*********************
import (
	"fmt"
	"sync"
)

func Squaring(num int, wt *sync.WaitGroup){
	defer wt.Done()
	fmt.Printf("%d squared equals %d\n", num, num*num)
	return
}

func main(){
	var wt sync.WaitGroup
	mass := []int{2, 4, 6, 8, 10}
	wt.Add(len(mass))
	for _, el := range mass{
		go Squaring(el, &wt)
	}
	wt.Wait()
}

//**********************Вариант 2*********************
// import (
// 	"fmt"
// )

// func Squaring(num int, ch chan int){
// 	fmt.Printf("%d squared equals %d\n", num, num*num)
// 	ch <- 1
// 	return 
// }

// func main(){
// 	mass := []int{2, 4, 6, 8, 10}
// 	ch := make(chan int)
// 	counter := len(mass)
// 	for _, el := range mass{
// 		go Squaring(el, ch)
// 	}
// 	for{
// 		counter = counter - (<-ch)
// 		if counter == 0{
// 			close(ch)
// 			return
// 		}
// 	}
// }