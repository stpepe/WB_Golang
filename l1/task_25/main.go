package main

import (
	"fmt"
	"time"
)

func Sleep(tm time.Time){
	for {
		if time.Now().After(tm){
			return
		}
	}
}

func main(){
	var body string
	fmt.Println(`Введите время для функции Sleep в формате: 3X, где Х - ns, us, µs, ms, s, m, h`)
	fmt.Scan(&body)
	tm, err := time.ParseDuration(body)
	if err != nil{
		fmt.Println("Неправильный формат времени!")
		return
	}
	fmt.Println(time.Now())
	Sleep(time.Now().Add(tm))
	fmt.Println(time.Now())
}