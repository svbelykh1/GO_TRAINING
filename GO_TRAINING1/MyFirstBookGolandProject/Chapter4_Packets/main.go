package main

import (
	"fmt"
	"log"

	gr "github.com/svbelykh1/GO_TRAINING/GO_TRAINING1/MyFirstBookGolandProject/internal/greeting"
	deu "github.com/svbelykh1/GO_TRAINING/GO_TRAINING1/MyFirstBookGolandProject/internal/greeting/deutsch"
	kb "github.com/svbelykh1/GO_TRAINING/GO_TRAINING1/MyFirstBookGolandProject/internal/keyboard"
)

func main() {

	//Вынесли функции в пакеты и тут их вызываем

	gr.Hi()
	deu.Hallo()
	fmt.Println("Введите результат: ")
	input, err := kb.Getfloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if input >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("Status:", status)

	//Консанты

	const TriangleSides = 3
	fmt.Println(TriangleSides)

}
