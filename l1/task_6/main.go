package main

import (
	"time"
	"context"
	"fmt"
)

func main(){
	// Остановка горутины вместе с основным потоком main
	go func(){
		for{}
	}()

	// Остановка горутины самостоятельно при помощи return
	go func(){
		fmt.Println("Go end - return")
		return
	}()

	// Остановка горутины при помощи ветвления case при закрытии канала
	stop := make(chan bool)
	go func(){
		for{
			select{
			default:
				time.Sleep(time.Second)
			case <-stop:
				fmt.Println("Go end - channel")
				return
			}
		}
	}()
	close(stop)

	// Остановка горутины при помощи ветвления case и контекста
	ctx, cancel := context.WithCancel(context.Background())
	go func(){
		for{
			select{
			default:
				time.Sleep(time.Second)
			case <-ctx.Done():
				fmt.Println("Go end - context")
				return
			}
		}
	}()
	cancel()

	time.Sleep(time.Second*1)
}