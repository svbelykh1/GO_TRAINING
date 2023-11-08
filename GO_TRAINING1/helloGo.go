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

	/*
		ЗАКЛЮЧЕНИЕ


		� Пакет представляет собой группу взаимосвязанных функций и других блоков кода.
		� Прежде чем использовать функции из пакета в файле Go, необходимо импортировать этот пакет.
		� Строка — последовательность байтов, обычно пред- ставляющих символы текста.
		� Руна представляет отдельный символ текста.
		� Два самых распространенных числовых типа Go — int (для хранения целых чисел) и float64 (для хранения чисел с плавающей точкой).
		� Тип bool используется для хранения логических зна- чений (true или false).
		� Переменная представляет собой блок памяти для хранения значения заданного типа.
		� Если переменной не присвоено значение, то она содер- жит нулевое значение для своего типа. Примеры нуле- вых значений: 0 для переменных int или float64, "" для строковых переменных.
		� Объявлениепеременнойможносовместитьсприсваива- нием ей значения при помощи короткого объявления переменной :=.
		� К переменной, функции или типу можно обращаться из кода других пакетов только в том случае, если ее имя начинается с буквы верхнего регистра.
		� Команда go fmt автоматически переформатирует исходные файлы по стандартам Go. Всегда выпол- няйте команду go fmt для любого кода, который вы собираетесь передавать другим разработчикам.
		� Команда go build компилирует исходный код Go в двоичный формат, который может выполняться компьютером.
		� Командаgorunкомпилируетивыполняетпрограмму без сохранения в исполняемом файле в текущем ката- логе.
	*/
}
