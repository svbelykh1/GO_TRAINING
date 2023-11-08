package main

import "fmt"

func main() {

	fmt.Println("hearth")
	//Типы
	/*
		fmt.Println(math.Floor(2.75))
		fmt.Println(strings.Title("я\n пошел\t \"за\" \t хлебом\\ "))
		fmt.Println('A')
		fmt.Println(reflect.TypeOf(34))
		fmt.Println(reflect.TypeOf(false))
		fmt.Println(reflect.TypeOf("34"))
		fmt.Println(reflect.TypeOf('A'))
	*/
	//Объявление переменных

	/*
		var names string
		names = "Petya"
		var family string = "Vasichkin"
		var otch = "Olenevich"
		ot := "Olenevich"
		var test_bool bool

		fmt.Println(names)
		fmt.Println(family)
		fmt.Println(otch)
		fmt.Println(ot)
		fmt.Println(test_bool)

	*/

	/*

		var myInt int = 2
		var myFloat float64 = 2.2
		//при сравнении будет ошибка, поэтому нужно преобразовать
		fmt.Println("myInt  > myFloat", myInt == int(myFloat))
		fmt.Println(reflect.TypeOf(myFloat))
		fmt.Println(reflect.TypeOf(int(myFloat)))
		fmt.Println(int(myFloat))
	*/

	/*
		var price int = 100
		fmt.Println("Price is", price, "dollars.")
		var taxRate float64 = 0.08
		var tax float64 = float64(price) * taxRate
		fmt.Println("Tax is", tax, "dollars.")
		var total float64 = float64(price) + tax
		fmt.Println("Total cost is", total, "dollars.")
		var availableFunds int = 120
		fmt.Println(availableFunds, "dollars available.")
		fmt.Println("Within budget?", total <= float64(availableFunds))
	*/

	//Форматировать код go fmt hello.go из консоли
	//Откомпилировать go build helloGo.go появиться исполняемый файл

}
