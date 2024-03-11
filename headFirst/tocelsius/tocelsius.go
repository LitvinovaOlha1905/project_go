// tocelsius преобразует температуру в градусах по Фаренгейту в градусы по Цельсию.
package main

import (
	"examples/headFirst/src/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()

	// Если функция вернула ошибку, программа сообщает об этом и завершается
	if err != nil {
		log.Fatal(err)
	}

	// Температура преобразуется к шкале Цельсия...
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius) // и выводится с двумя знаками после точки.
}
