package chapter2

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func chapter2() {
	fmt.Println("Chapter 2")
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
	///////////////////////////////////////////////////
	str := "Go r#ocks!"
	replacer := strings.NewReplacer("#", "o")
	fmt.Println(reflect.TypeOf(replacer))
	strNew := replacer.Replace(str)
	fmt.Println(strNew)

	str1 := "|__||__||__|"
	str2 := strings.NewReplacer("||", "__").Replace(str1)
	fmt.Println(str2)
	//  Т.е мы подключили пакет strings хотим применить репласе к строке str
	// у пакета strings есть метод NewReplacer, который содержит правило по которому делаем замену(кладем это в переменную replacer)  и у этой переменной есть метод Replace(получает на вход строку)

}
