package main

import (
	"fmt"
)

func main() {
	temps := map[int][]float32{}

	mass := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, el := range mass {
		group := int(el) / 10 * 10

		if temps[group] == nil {
			temps[group] = []float32{}
		}

		temps[group] = append(temps[group], el)
	}

	fmt.Println(temps)
}