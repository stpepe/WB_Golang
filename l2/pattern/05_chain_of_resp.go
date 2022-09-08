package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
import (
	"fmt"
	"crypto/sha256"
	"strings"
)

type Chain interface{
	execute([]string)
}

type AddToBlockchain struct{
	next Chain
}

func (a *AddToBlockchain) execute(blockchain []string){
	if len(blockchain) < 2{
		a.next.execute(blockchain)
		return
	}
	var sb strings.Builder
	sb.WriteString(blockchain[len(blockchain)-1])
	sb.WriteString(blockchain[len(blockchain)-2])
	blockchain[len(blockchain)-1] = sb.String()
	a.next.execute(blockchain)
}

type GenerateHash struct{
	
}

func (g *GenerateHash) execute(blockchain []string){
	hash := sha256.New()
	hash.Write([]byte(blockchain[len(blockchain)-1]))
	blockchain[len(blockchain)-1] = string(hash.Sum(nil))
	fmt.Printf("%x\n",blockchain)
	return 
}

func main(){
	blockchain := []string{}
	var el string
	generateHash := &GenerateHash{}
	addToBlockchain := &AddToBlockchain{next: generateHash}
	for{
		fmt.Println("Введите слово, которое необходимо добавить в блокчейн или exit, чтобы закончить:")
		fmt.Scan(&el)
		if el == "exit"{
			return
		}
		blockchain = append(blockchain, string(el))
		addToBlockchain.execute(blockchain)
	}
}



