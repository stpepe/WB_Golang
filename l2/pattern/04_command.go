package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

import (
	"fmt"
)

type Device interface{
	switch_state()
	forward()
	back()
}

type Command interface{
	execute()
}

type SwitchStateCommand struct{
	Device
}

func (s *SwitchStateCommand) execute(){
	s.switch_state()
}

type ForwardCommand struct{
	Device
}

func (f *ForwardCommand) execute(){
	f.forward()
}

type BackCommand struct{
	Device
}

func (b *BackCommand) execute(){
	b.back()
}

type Button struct{
	Command
}

func (b *Button) press(){
	b.execute()
}

type TV struct{
	state bool
	channel int
}

func (tv *TV) switch_state(){
	if tv.state{
		tv.state = false
		fmt.Println("Телевизор выключен")
	} else {
		tv.state = true
		fmt.Println("Телевизор включен")
	}

}

func (tv *TV) forward(){
	if tv.state{
		if tv.channel < 99{
			tv.channel += 1
		} else {
			tv.channel = 0
		}
		fmt.Println("Вы на канале номер:", tv.channel)
	} else {
		fmt.Println("Нельзя переключать каналы у выключенного телевизора")
	}
}

func (tv *TV) back(){
	if tv.state{
		if tv.channel > 0{
			tv.channel -= 1
		} else {
			tv.channel = 99
		}
		fmt.Println("Вы на канале номер:", tv.channel)
	} else {
		fmt.Println("Нельзя переключать каналы у выключенного телевизора")
	}
}

func main(){
	tv := &TV{state: false, channel: 0}
	ButtonForward := ForwardCommand{Device: tv}
	ButtonBack := BackCommand{Device: tv}
	ButtonState := SwitchStateCommand{Device: tv}

	ButtonForward.execute()
	ButtonBack.execute()

	ButtonState.execute()
	ButtonForward.execute()
	ButtonForward.execute()
	ButtonForward.execute()
	ButtonBack.execute()
	ButtonState.execute()

}
