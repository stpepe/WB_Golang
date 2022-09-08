package main	

import (
	"fmt"
	"sync"
)

type Counter struct{
	mut sync.Mutex
	value int 
} 

func NewCounter() *Counter{
	v := Counter{
		value: 0,
	}
	return &v
}

func Inc(counter *Counter){
	counter.mut.Lock()
	defer counter.mut.Unlock()
	counter.value++
	return
}

func main(){
	counter := NewCounter()
	var wt sync.WaitGroup 

	for i:=0; i<14; i++{
		wt.Add(1)
		go func(){
			defer wt.Done()
			Inc(counter)
		}()
	}
	
	wt.Wait()
	fmt.Println(counter.value)

}