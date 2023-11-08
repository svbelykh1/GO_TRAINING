package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Можно посмотеть как работает strconv и Replace
func main() {

	//Размер файла
	/*
		fileinfo, err := os.Stat("my.txt")
		if err != nil {
			log.Fatal(err)
		}
		size := fileinfo.Size()
		fmt.Println(size)
	*/

	//получаем ввод с клавиатуры в переменную input

	fmt.Print("Введите любую строку  ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	//1) Вариант использовать Replace
	//replacer := strings.NewReplacer("\n", "")
	//rInput := replacer.Replace(input)
	//2) использовать TrimSpace у пакета strings, которая удалит все пробеоы табуляции и переносы строк

	rInput := strings.TrimSpace(input)
	//fmt.Println(rInput)
	toRInput, err := strconv.ParseFloat(rInput, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if toRInput >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println(status)

}
