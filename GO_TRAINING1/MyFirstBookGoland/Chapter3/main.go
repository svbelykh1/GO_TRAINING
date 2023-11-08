package main

import "fmt"

var metersPerLiter float64

func main() {

	//var width, height, area float64
	//width = 4.2
	//height = 7.7
	//area = width * height
	//fmt.Printf("%f Потребуется литров - формотируемый вывод с помощью printf\n", area/10)
	//fmt.Println(area / 10)
	//fmt.Printf("Форматированная строка %0.2f\n", 1.0/3.0) // функция Printf выводит строку
	//fmt.Printf(fmt.Sprintf("Форматированная строка %0.2f\n", 1.0/3.0)) // возвращает строку

	//fmt.Printf("ЭТО строка в котрое вставлено одно значение %v и второе значение %v, а это их произведение %#v ну и перевод строки %#v", width, height, width*height, "\n")

	//fmt.Printf("%20s | %s", "product", "cost in cents\n")
	//fmt.Println("------------------------------------------------")
	//fmt.Printf("%20s | %2d\n", "туалетная бумага", 50)
	//fmt.Printf("%20s | %2d\n", "ЧЕХОЛ ДЛЯ БУМАГИ", 5)

	//Форматирование дробого числа к примеру происходит так ...
	//float_ := 121.11110000791
	//fmt.Printf("Форматирование дробого числа к примеру происходит так ... доаустим есть число %#v\n", float_)
	//fmt.Printf("для его форматирование указывается %%0.2f, где 0 - это минимальная ширина всего числа, а 2 ширина дробной части\n")
	//fmt.Printf("%s - %0.2f\n", "результат", float_)

	// ---------------------------------------------------------------------------------------------------
	//функции
	//sayHello("Привет,", "мир", 5)
	metersPerLiter = 7
	fmt.Printf("%.2f литров требуется для покраски\n", paintNeed(3.7, 4.4))
}
func sayHello(var1 string, var2 string, var3 int) {
	for x := 1; x <= var3; x++ {
		println(var1, var2, x)
	}
}

func paintNeed(height, width float64) float64 {
	area := height * width
	return area / metersPerLiter
}
