package main

import (
	"fmt"
	"log"

	gr "github.com/svbelykh1/GO_TRAINING/GO_TRAINING1/MyFirstBookGolandProject/internal/greeting"
	kb "github.com/svbelykh1/GO_TRAINING/GO_TRAINING1/MyFirstBookGolandProject/internal/keyboard"
)

func main() {
	gr.Hi()
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
}
