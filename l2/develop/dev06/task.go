package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
)

type Flags struct {
	F int
	D string
	S bool
}

func main() {
	flagF := flag.Int("f", 0, "")
	flagD := flag.String("d", " ", "")
	flagS := flag.Bool("s", false, "")

	flag.Parse()

	flags := Flags{
		F: *flagF,
		D: *flagD,
		S: *flagS,
	}

	fmt.Println("Введите название файла или просто enter, если хотите вводить строки в консоль:")
	var fileName string = ""
	fmt.Scanln(&fileName)

	lines := []string{}
	if fileName != "" {
		bytes, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatalln(err)
		}
		lines = strings.Split(string(bytes), "\n")
	} else {
		fmt.Println("Введите строки:")
		for {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			if scanner.Text() == ""{
				break
			}
			lines = append(lines, scanner.Text()) 
		}	
	}

	for _, line := range lines {
		cols := strings.Split(line, flags.D)

		if flags.F > 0 && flags.F <= len(cols) {
			fmt.Println(cols[flags.F-1])
		} else if !flags.S {
			fmt.Println(line)
		}
	}
}
