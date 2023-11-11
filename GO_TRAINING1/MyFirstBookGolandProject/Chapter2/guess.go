// игра, в которой игрок должен угадать случайное число
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	target := rand.Intn(100) + 1
	status := false
	for i := 1; i < 9; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Введите число: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi((input))
		if err != nil {
			log.Fatal(err)
		}
		if guess > target {
			fmt.Println("Указанное число " + strconv.Itoa(guess) + " БОЛЬШЕ загаданного")
			fmt.Println("У тебя осталось  " + strconv.Itoa(8-i) + "  попыток")
		} else if guess < target {
			fmt.Println("Указанное число " + strconv.Itoa(guess) + " МЕНЬШЕ загаданного")
			fmt.Println("У тебя осталось  " + strconv.Itoa(8-i) + "  попыток")
		} else {
			fmt.Println("ПОЗДРАВЛЯЮ !!! Ты угадал, загаданное число " + strconv.Itoa(target))
			status = true
			break
		}

	}
	if !status {
		fmt.Println("Ты не смог угадать, загаданное число :(" + strconv.Itoa(target))
	}
}
