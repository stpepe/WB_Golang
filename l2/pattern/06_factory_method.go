package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
import (
	"fmt"
)

const (
	sportCar = "sport"
	truckCar = "truck"
)

type Car interface{
	GetName()
	MaxSpeed()
	TypeCar()
}

type Truck struct{
	name string
	maxSpeed float32
	truckType string
}

func CreateTruck(name, truckType string, maxSpeed float32) *Truck {
	return &Truck{
		name: name,
		truckType: truckType,
		maxSpeed: maxSpeed,
	}
}

func (t *Truck) GetName(){
	fmt.Println(t.name)
}

func (t *Truck) MaxSpeed(){
	fmt.Println(t.maxSpeed)
}

func (t *Truck) TypeCar(){
	fmt.Println(t.truckType)
}

type SportCar struct{
	name string
	maxSpeed float32
	carType string
}

func CreateSportCar(name, carType string, maxSpeed float32) *SportCar {
	return &SportCar{
		name: name,
		carType: carType,
		maxSpeed: maxSpeed,
	}
}

func (t *SportCar) GetName(){
	fmt.Println(t.name)
}

func (t *SportCar) MaxSpeed(){
	fmt.Println(t.maxSpeed)
}

func (t *SportCar) TypeCar(){
	fmt.Println(t.carType)
}

func CreateCar(str string) Car {
	switch str {
	default:
		fmt.Println("Такой машины нет!")
		return nil
	case sportCar:
		return CreateSportCar("Ferrari", sportCar, 300)
	case truckCar:
		return CreateTruck("Man", truckCar, 100)	
	}
}

func main(){
	fmt.Println("Введите машину, которую хотите создать (sport, truck):")
	var t string
	fmt.Scan(&t)
	car := CreateCar(t)
	if car == nil{
		return
	}
	car.GetName()
	car.MaxSpeed()
	car.TypeCar()
}