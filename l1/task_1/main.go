package main

import (
	"fmt"
)

type Human struct {
	name string
	sec_name string
	age int
}

type Action struct {
	Human
}

func CreateHuman(name string, sec_name string, age int) Human{
	h := Human{
		name: name,
		sec_name: sec_name,
		age: age,
	}

	return h
}

func (h *Human) Walking(){
	fmt.Printf("%s is walking\n", h.name)
	return
}

func main(){
	man := CreateHuman("Dude", "Just", 43)
	act := Action{Human: man}
	act.Walking() 
}