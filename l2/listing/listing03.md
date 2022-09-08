Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

В случае с интерфейсом, мы сравниваем и значение и тип. У nil тип nil, поэтому err не равно nil.
Чтобы преобразовать его к типу *os.PathError воспользуемся (*os.PathError)(nil)

```
