package main

import (
	"fmt"
	// "math/rand"
)

func quicksort(a []int) []int {
    if len(a) < 2 {
        return a
    }
     
    left, right := 0, len(a)-1
     
	pivot := len(a) / 2

    // pivot := rand.Int() % len(a)
     
    a[pivot], a[right] = a[right], a[pivot]
     
    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }
     
    a[left], a[right] = a[right], a[left]
     
    quicksort(a[:left])
    quicksort(a[left+1:])
     
    return a
}

func main(){
	unsorted := []int{1, 10, 14, 12, 2, 111}
	sorted := quicksort(unsorted)
	fmt.Println(sorted)
}