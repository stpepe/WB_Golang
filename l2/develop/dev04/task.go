package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
)

func hashCalc(word string) uint32 {
	var hash uint32 = 0
	for _, v := range []rune(word){
		hash += uint32(v) 
	}
	return hash
}

func searchSets(mass *[]string) map[string][]string{
	tmpDictionary := make(map[uint32][]string)
	for _, v := range *mass{
		vHash := hashCalc(v)
		if _, ok := tmpDictionary[vHash]; !ok{
			arr := []string{v}
			tmpDictionary[vHash] = arr
		} else {
			tmpDictionary[vHash] = append(tmpDictionary[vHash], v)
		}
	}
	resultDictionary := make(map[string][]string)
	for _, value := range tmpDictionary{
		resultDictionary[value[0]] = value
	}
	return resultDictionary
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "лоло", "олло"}
	fmt.Println(searchSets(&words))
}
