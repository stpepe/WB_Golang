package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"sort"
	"reflect"
)

var flags = map[string]int{
	"-k": 1, 
	"-n": 0,
	"-r": 0,
	"-u": 0,   
}

var flagsPrior = map[string]int{
	"-k": 0, 
	"-n": 1,
	"-r": 1,
	"-u": 2,   
}


func formatInput (input string)(string, map[string]int){
	splitedInput := strings.Split(input, " ")
	fileName := splitedInput[0]
	splitedInput = splitedInput[1:]
	flagsMap := make(map[string]int)
	if len(splitedInput) > 0{
		for i, value := range splitedInput{
			if _, ok := flags[value]; ok{
				if flags[value] == 1{
					_, okk := strconv.Atoi(splitedInput[i+1])
					if okk != nil{
						log.Fatal("Неверный формат флагов!")
					}
					flagsMap[value], _ = strconv.Atoi(splitedInput[i+1])
				} else {
					flagsMap[value] = 0
				}
			} else {
				log.Fatal("Введен несуществующий флаг(и)!")
			}
		} 
	}
	return fileName, flagsMap
}

type sortObject struct{
	columnNum int
	fileName string
	fileData map[int][]string
	sorter
}

func (s *sortObject) startSort(){
	s.sorter.execute(s)
	file, err := os.Create(s.fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	for _, value := range s.fileData{
		var str string
		str = strings.Join(value, " ")
		str += "\n"
		_, err2 := file.WriteString(str)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	

}

type sorter interface{
	execute(*sortObject)
}

type kSorter struct{
	next sorter
	last sorter
}

func (k *kSorter) execute(s *sortObject){
	column := make(map[int][]string)
	tmpElement := []string{} 
	for _, value := range s.fileData{
		for i, v := range value{
			if i == s.columnNum-1{
				tmpElement = append(tmpElement, v)
			}
		}
	}
	column[s.columnNum-1] = tmpElement
	tmpFileData := make(map[int][]string)
	tmpFileData, s.fileData = s.fileData, column
	fmt.Println(s.fileData)
	k.next.execute(s)
	iterator := 0
	for key := range tmpFileData{
		tmpFileData[key][s.columnNum-1] = s.fileData[s.columnNum-1][iterator]
		iterator++
	}
	s.fileData = tmpFileData
	if k.last != nil{
		k.last.execute(s)
	}
}

type stringSorter struct{
	next sorter
}

func (st *stringSorter) execute(s *sortObject){
	for key, value := range s.fileData{
		sort.Strings(value)
		s.fileData[key] = value
	}
	if st.next != nil{
		st.next.execute(s)
	}
}

type numSorter struct{
	next sorter
}

func (n *numSorter) execute(s *sortObject){
	for key, value := range s.fileData{
		tmpIntElements := []int{}
		for _, v := range value{
			tmpElement, ok := strconv.Atoi(v)
			if ok != nil{
				log.Fatal("В файле для сортировки по числам не должно быть нечисловых значений!")
			}
			tmpIntElements = append(tmpIntElements, tmpElement)
		}
		sort.Ints(tmpIntElements)
		tmpStrElements := []string{}
		for _, v := range tmpIntElements{
			tmpStrElements = append(tmpStrElements, strconv.Itoa(v))
		}
		s.fileData[key] = tmpStrElements
	}
	if n.next != nil{
		n.next.execute(s)
	}
}

type reverseSorter struct{
	next sorter
}

func (r *reverseSorter) execute(s *sortObject){
	for key, value := range s.fileData{
		revValue := []string{}
		for i := len(value) - 1; i>= 0; i-- {
    		revValue = append(revValue, value[i])
		}
		s.fileData[key] = revValue
	}
	if r.next != nil{
		r.next.execute(s)
	}
}

type uSorter struct{
	next sorter
}

func (u *uSorter) execute(s *sortObject){
	for key1, value1 := range s.fileData{
		for key2, value2 := range s.fileData{
			if reflect.DeepEqual(value1, value2) && key1 != key2{
				delete(s.fileData, key2)
			}
		}
	}
}

type sortBuilder struct{

}

func (b *sortBuilder) buildSort(fileName string, flagsMap map[string]int) *sortObject {
	fileData := make(map[int][]string)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		fileData[i] = strings.Split(scanner.Text(), " ")
	}
		
	sortObj := &sortObject{
		fileName: fileName,
		fileData: fileData,
	}

	var ok bool
	var columnNum int
	if columnNum, ok = flagsMap["-k"]; ok{
		k := &kSorter{}
		sortObj.columnNum = columnNum
		sortObj.sorter = k
		if _, ok = flagsMap["-u"]; ok{
			k.last = &uSorter{}
		} 
		if _, ok = flagsMap["-n"]; ok{
			n := &numSorter{}
			k.next = n
			if _, ok = flagsMap["-r"]; ok{
				r := &reverseSorter{}
				n.next = r
			}
		} else {
			s := &stringSorter{}
			k.next = s
			if _, ok = flagsMap["-r"]; ok{
				r := &reverseSorter{}
				s.next = r
			}
		}
		
	} else {
		if _, ok = flagsMap["-n"]; ok{
			n := &numSorter{}
			sortObj.sorter = n
			if _, ok = flagsMap["-r"]; ok{
				r := &reverseSorter{}
				n.next = r
				if _, ok = flagsMap["-u"]; ok{
					u := &uSorter{}
					r.next = u
				}
			}
		} else {
			s := &stringSorter{}
			sortObj.sorter = s
			if _, ok = flagsMap["-r"]; ok{
				r := &reverseSorter{}
				s.next = r
				if _, ok = flagsMap["-u"]; ok{
					u := &uSorter{}
					r.next = u
				}
			} else {
				u := &uSorter{}
				s.next = u
			}
		}
	}
	return sortObj
}


func main() {
	fmt.Println("Введите имя файла и флаги:")
	stdScanner := bufio.NewScanner(os.Stdin)
	stdScanner.Scan()
	builder := &sortBuilder{}
	sortObj := builder.buildSort(formatInput(stdScanner.Text()))
	sortObj.startSort()
	fmt.Println("Сортировка окончена!")
}
