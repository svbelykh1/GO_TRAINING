package keyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Getfloat() (float64, error) {
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
