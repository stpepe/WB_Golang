package main

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"fmt"
	"time"
	"os"
	"github.com/beevik/ntp"
)

const (
	timeFormat = "15:04:05"
)

func main() {
	t, err := ntp.Time("ge.pool.ntp.org")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Время через ntp:", t.Format(timeFormat))
	fmt.Println("Время через Now():", time.Now().Format(timeFormat))
}
