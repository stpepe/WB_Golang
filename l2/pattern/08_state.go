package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import (
	"fmt"
)

type State interface{
	Go()bool
	Wait()bool
	Prepare()bool
}

type RedState struct{

}

func (r *RedState) Go()bool{
	fmt.Println("Вы не можете идти, свет красный")
	return false
}

func (r *RedState) Wait()bool{
	fmt.Println("Ожидайте желтый свет")
	return true
}

func (r *RedState) Prepare()bool{
	fmt.Println("Вы не можете готовиться, свет красный")
	return false
}

type GreenState struct{

}

func (r *GreenState) Go()bool{
	fmt.Println("Правильно, идите!")
	return true
}

func (r *GreenState) Wait()bool{
	fmt.Println("Не нужно ждать, горит зеленый!")
	return false
}

func (r *GreenState) Prepare()bool{
	fmt.Println("Не нужно готовиться, горит зеленый!")
	return false
}

type YellowState struct{

}

func (r *YellowState) Go()bool{
	fmt.Println("Нельзя идти, горит желтый!")
	return false
}

func (r *YellowState) Wait()bool{
	fmt.Println("Не нужно ждать, уже желтый!")
	return false
}

func (r *YellowState) Prepare()bool{
	fmt.Println("Правильно, готовьтесь к переходу!")
	return true
}

type TrafficLight struct{
	red State
	yellow State
	green State

	current State
}

func (t *TrafficLight) SetCurrentState(state State){
	t.current = state
}

func (t *TrafficLight) Go(){
	if t.current.Go(){
		t.SetCurrentState(t.red)
		return
	}
}

func (t *TrafficLight) Wait(){
	if t.current.Wait(){
		t.SetCurrentState(t.yellow)
		return
	}
}

func (t *TrafficLight) Prepare(){
	if t.current.Prepare(){
		t.SetCurrentState(t.green)
		return
	}
}

func main(){
	red := &RedState{}
	green := &GreenState{}
	yellow := &YellowState{}

	light := &TrafficLight{
		red: red,
		green: green,
		yellow: yellow,
		current: red,
	}

	light.Go()
	light.Prepare()
	light.Wait()
	light.Prepare()
	light.Go()
}