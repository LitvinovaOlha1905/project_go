// tocelsius преобразует температуру в градусах по Фаренгейту в градусы по Цельсию.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)   // Создать «буферизованное средство чтения» для получе-ния текста с клавиатуры.
	input, err := reader.ReadString('\n') // Возвращает текст, введенный пользователем до нажатия клавиши Enter.
	// Сообщить о любых ошибках при преобразовании.
	if err != nil {
		return 0, err // ...если при чтении ввода произошла ошибка, она бу-дет возвращена функцией.
	}

	input = strings.TrimSpace(input)             // Удаляем символ новой строки из входной строки.
	number, err := strconv.ParseFloat(input, 64) // Преобразовать строку в значение float64.
	if err != nil {
		return 0, err // Также возвращаются ошибки, которые могут возникнуть при преобразовании строки в float64.
	}
	return number, nil
}

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := getFloat() // Вызываем getFloat для получения температуры.

	// Если функция вернула ошибку, программа сообщает об этом и завершается
	if err != nil {
		log.Fatal(err)
	}

	// Температура преобразуется к шкале Цельсия...
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius) // и выводится с двумя знаками после точки.
}
