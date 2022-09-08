package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"log"
	"io/ioutil"
	// "os"
	"regexp"
	"strings"
)

type Flags struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
}

func main() {
	flagA := flag.Int("A", 0, "")
	flagB := flag.Int("B", 0, "")
	flagC := flag.Int("C", 0, "")
	flagc := flag.Bool("c", false, "")
	flagi := flag.Bool("i", false, "")
	flagv := flag.Bool("v", false, "")
	flagF := flag.Bool("F", false, "")
	flagn := flag.Bool("n", false, "")

	flag.Parse()

	f := Flags{
		A: *flagA,
		B: *flagB,
		C: *flagC,
		c: *flagc,
		i: *flagi,
		v: *flagv,
		F: *flagF,
		n: *flagn,
	}

	fmt.Println("Введите имя файла")
	var fileName string 
	fmt.Scan(&fileName)

	fmt.Println("Введите регулярное выражение:")
	var reg string
	fmt.Scan(&reg)

	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	count := 0

	query := reg

	after := f.A
	before := f.B

	if f.C > 0 {
		after = f.C
		before = f.C
	}

	if f.i {
		query = strings.ToLower(query)
	}

	for i, v := range lines {
		isPrint := false

		if f.i {
			v = strings.ToLower(v)
		}

		if f.F {
			isPrint = query == v
		} else {
			isPrint, _ = regexp.Match(query, []byte(v))
		}

		if f.v {
			isPrint = !isPrint
		}

		if isPrint {
			count++

			if !f.c {
				for j := 1; j <= before; j++ {
					k := i - j
					if k >= 0 {
						if f.n {
							fmt.Printf("%v: %v\n", k+1, lines[k])
						} else {
							fmt.Println(lines[k])
						}
					}
				}

				if f.n {
					fmt.Printf("%v: %v\n", i+1, v)
				} else {
					fmt.Println(v)
				}

				for j := 1; j <= after; j++ {
					k := i + j
					if k <= len(lines) {
						if f.n {
							fmt.Printf("%v: %v\n", k+1, lines[k])
						} else {
							fmt.Println(lines[k])
						}
					}
				}
			}
		}
	}

	if f.c {
		fmt.Println(count)
	}
}
