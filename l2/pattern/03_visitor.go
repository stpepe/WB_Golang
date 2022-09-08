package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

import (
	"fmt"
)

type Accepter interface{
	accept(v Visitor)
}

type Square struct{
	length int
}

func (s *Square) accept(v Visitor){
	fmt.Println(v.VisitForSquare(s))
}

type Circle struct{
	radius float32
}

func (c *Circle) accept(v Visitor){
	fmt.Println(v.VisitForCircle(c))
}

type Visitor interface{
	VisitForSquare(s *Square) int
	VisitForCircle(c *Circle) float32
}

type AreaCalculator struct{

}

func (a *AreaCalculator) VisitForSquare(s *Square) int {
	return s.length * s.length
}

func (a *AreaCalculator) VisitForCircle(c *Circle) float32 {
	return c.radius * c.radius * 3.14
}

func main(){
	areaCalculator := &AreaCalculator{}
	circle := &Circle{radius: 5}
	square := &Square{length: 5}
	circle.accept(areaCalculator)
	square.accept(areaCalculator)
}