package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

func Writer(mut *sync.Mutex, mapp *[]int, wt *sync.WaitGroup){
	defer wt.Done()
	var dur int = 0
	for i:=0; i<5; i++{
		dur = rand.Intn(500)
		time.Sleep(time.Millisecond * time.Duration(dur))
		mut.Lock()
		*mapp = append(*mapp, i)
		mut.Unlock()
	}
	return
}

func main(){
	mapp := []int{}
	var mut sync.Mutex
	var wt sync.WaitGroup

	for i:=0; i<3; i++{
		wt.Add(1)
		go Writer(&mut, &mapp, &wt)
	}

	wt.Wait()
	fmt.Println(mapp)

}