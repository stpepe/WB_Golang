package main

import (
	"os"
	"os/signal"
	"syscall"
	"math/rand"
	"sync"
	"time"
	"fmt"
)

func Publisher(ch, req chan int, stop chan bool){
	for{
		select{
		case <-req:
			ch <- rand.Intn(1000)
		case <-stop:
			return
		}
	}
}

func Worker(ch, req chan int, stop chan bool){
	i:=0
	for{
		select{
		default:
			i++
			req <- i
			fmt.Println(<-ch)
			time.Sleep(time.Second)
		case <-stop:
			return
		}
	}
}

func main(){
	ch := make(chan int)
	req := make(chan int)
	stop := make(chan bool)
	ex_chan := make(chan os.Signal)
	var wt sync.WaitGroup
	signal.Notify(ex_chan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Print("Введите целое число воркеров: ")
	var workers int
	_, err := fmt.Fscan(os.Stdin, &workers)
	if err != nil {
		panic(err)
	}
	go func(){
		wt.Add(1)
		Publisher(ch, req, stop)
		wt.Done()
	}()

	for i:=0; i<workers; i++{
		go func(){
			wt.Add(1)
			Worker(ch, req, stop)
			wt.Done()
		}()
	}

	<-ex_chan
	close(stop)
	wt.Wait()
	close(ch)
	close(req)
	close(ex_chan)
	fmt.Println("Exit")
}
