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
	var status string
	fmt.Print("Введите число: \n")
	input, err := getfloat()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(input)

	if input >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println(status)
}

func getfloat() (float64, error) {
	//ожидание и обработка ввода с клавиатуры
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	//обработка введенного значения, использовать TrimSpace у пакета strings, которая удалит все пробеоы табуляции и переносы строк
	inp = strings.TrimSpace(inp)
	//
	toFloatInput, err := strconv.ParseFloat(inp, 64)
	if err != nil {
		return 0, err
	}
	return toFloatInput, nil
}
