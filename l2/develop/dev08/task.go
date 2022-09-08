package main

/*
=== Взаимодействие с ОС ===
Необходимо реализовать собственный шелл
встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах
Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

func CmdCd(c []string) {
	myDirr, err := os.Getwd()
	if err != nil {
		return
	}

	if (len(c) > 1 && c[1] == "~") || len(c) == 1 {
		home, err := os.UserHomeDir()
		if err != nil {
			return
		}
		os.Chdir(filepath.Join(home))
	} else if len(c) > 1 {
		os.Chdir(filepath.Join(myDirr, c[1]))
	}
}

func CmdPwd(c []string) {
	myDirr, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(myDirr)
}

func CmdEcho(c []string) {
	for _, e := range c[1:] {
		fmt.Print(e, " ")
	}
	fmt.Println()
}

func CmdKill(c []string) {
	for _, e := range c[1:] {
		pid, err := strconv.Atoi(e)
		if err != nil {
			continue
		}
		pro, err := os.FindProcess(pid)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
		err = pro.Kill()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
	}
}

func CmdPs(c []string) {
	processes, err := ps.Processes()
	if err != nil {
		return
	}
	fmt.Printf("%s\t%s\n", "PID", "COMMAND")
	for _, proc := range processes {
		fmt.Printf("%d\t%s\n", proc.Pid(), proc.Executable())
	}
}

func CmdQuit(c []string) {
	os.Exit(0)
}

func main() {
	myDirr, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Print(myDirr, "> ")

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		for _, cmd := range strings.Split(reader.Text(), ";") {
			slice := strings.Fields(cmd)
			if len(slice) == 0 {
				myDirr, err := os.Getwd()
				if err != nil {
					return
				}
				fmt.Print(myDirr, "> ")
				continue
			}

			definedCommands := map[string]func(c []string){
				"cd":   CmdCd,
				"pwd":  CmdPwd,
				"echo": CmdEcho,
				"kill": CmdKill,
				"quit": CmdQuit,
				"ps":   CmdPs,
			}
			if f, ok := definedCommands[slice[0]]; ok {
				f(slice)
			} else {
				var cmd *exec.Cmd
				if len(slice) > 1 {
					cmd = exec.Command(slice[0], slice[1:]...)
				} else {
					cmd = exec.Command(slice[0])
				}
				res, _ := cmd.Output()
				fmt.Print(string(res))
			}
		}
		myDirr, err := os.Getwd()
		if err != nil {
			return
		}
		fmt.Print(myDirr, "> ")
	}
}
