Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

В 1 функции мы возвращаем переменную х, но не через return.
Поэтому defer, вызов которой следует после вызова return добавил
к значению ч еще единицу

Во 2 функции мы возвращаем переменную х через return, из-за чего
возвращается значение х в момент return. А как мы знаем, функция
defer идет после вызова return, слеловательно defer не успел добавить
единицу.

```
