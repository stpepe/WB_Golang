package main

import (
	"fmt"
	"bufio"
	"os"
)

type OldMessage struct{
	message []byte
}

func CreateOldMessage(str string) *OldMessage{
	bstr := []byte(str)
	m := OldMessage{
		message: bstr,
	}
	return &m
}

func (m *OldMessage) Send() []byte {
	return m.message
}

type ModerMessage struct{
	message []rune	
}

func (m *ModerMessage) Receive(str []rune){
	m.message = str
	fmt.Println("Полученное сообщение: ",string(str))
}

type MessageAdapter struct{
	oldmsg *OldMessage
}

func (m *MessageAdapter) AdaptSend() []rune {
	adapt_message := []rune(string(m.oldmsg.Send()))
	return adapt_message
}

func main(){
	var body string
	fmt.Println("Введите сообщение:")
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    body = scanner.Text()
	old := CreateOldMessage(body)

	modern := &ModerMessage{}
	adapter := &MessageAdapter{oldmsg: old}
	// modern.Receive(old.Send()) так не работает, хотя нам нужно чтобы работало
	modern.Receive(adapter.AdaptSend())
}