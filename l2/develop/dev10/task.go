package main

/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123
Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"errors"
)

var ErrNotEnoughArguments error = errors.New("Not enough arguments.")
var ErrIncorrectArgument error = errors.New("Incorrect argument.")
var ErrCanNotConvertTimeString error = errors.New("Can not convert time string.")
var ErrConnectionClosed error = errors.New("Connection closed.")

func ConvertTimeString(str string) (time.Duration, error) {
	letter := str[len(str)-1:]
	numberStr := str[:len(str)-1]
	switch letter {
	case "s":
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return *new(time.Duration), err
		}

		return time.Second * time.Duration(number), nil
	}

	return *new(time.Duration), ErrCanNotConvertTimeString
}

type Args struct {
	url     string
	port    int
	timeout time.Duration
}

func CreateArgs(a []string) (Args, error) {
	strArgs := a[1:]

	args := Args{
		timeout: time.Second * 10,
		//timeout: time.Millisecond,
	}

	if len(strArgs) < 2 {
		return args, ErrNotEnoughArguments
	}

	mainArgs := strArgs[len(strArgs)-2:]

	args.url = mainArgs[0]

	port, err := strconv.Atoi(mainArgs[1])
	if err != nil {
		return args, err
	}
	args.port = port

	if len(strArgs) > 2 {
		otherArgs := strArgs[:len(strArgs)-2]

		for _, v := range otherArgs {
			arr := strings.Split(v, "=")

			if len(arr) != 2 {
				return args, ErrIncorrectArgument
			}

			switch arr[0] {
			default:
				return args, ErrIncorrectArgument
			case "--timeout":
				t, err := ConvertTimeString(arr[1])
				if err != nil {
					return args, ErrIncorrectArgument
				}
				args.timeout = t
			}
		}
	}

	return args, nil
}

func (v *Args) GetAddr() string {
	return v.url + ":" + strconv.Itoa(v.port)
}

func (v *Args) GetURL() string {
	return v.url
}

func main() {
	args, err := CreateArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), args.timeout)
	defer cancel()

	d.LocalAddr = nil

	fmt.Printf("Trying %v...\n", args.GetURL())
	conn, err := d.DialContext(ctx, "tcp", args.GetAddr())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("Connected to %v.\n", args.GetURL())

	for {
		var str string
		_, err := fmt.Fscan(os.Stdin, &str)
		if err != nil {
			if err == io.EOF {
				conn.Close()
				fmt.Println(ErrConnectionClosed)
				return
			}
			panic(err)
		}

		_, err = fmt.Fprintf(conn, fmt.Sprintf("%v\r\n\r\n", str))
		if err != nil {
			fmt.Println(ErrConnectionClosed)
			return
		}

		answer, err := io.ReadAll(conn)
		if err != nil {
			fmt.Println(ErrConnectionClosed)
			return
		}
		fmt.Println(string(answer))
	}
}
