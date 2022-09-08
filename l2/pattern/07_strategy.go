package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/
import (
	"fmt"
	"bufio"
	"os"
)

type Sender interface{
	Send(string)
}

type Message struct{
	msg string
}

func (m *Message) SendMessage(s Sender){
	s.Send(m.msg)
}

type MailRu struct{

}

func (m *MailRu) Send (str string){
	fmt.Printf("Отправлено через mail.ru: %s\n", str)
} 

type Telegram struct{

}

func (t *Telegram) Send (str string){
	fmt.Printf("Отправлено через Telegram: %s\n", str)
} 

type WhatsApp struct{

}

func (w *WhatsApp) Send (str string){
	fmt.Printf("Отправлено через WhatsApp: %s\n", str)
} 


func main(){
	fmt.Println("Введите сообщение, которое необходимо отправить:")
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    body := scanner.Text()
	msg := &Message{msg: body}
	mail := &MailRu{}
	tg := &Telegram{}
	ws := &WhatsApp{}

	msg.SendMessage(mail)
	msg.SendMessage(tg)
	msg.SendMessage(ws)
}