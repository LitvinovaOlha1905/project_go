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
	fmt.Println("Enter a grade: ") // Запрашиваем у поль-зователя значение.

	// Вызываем getFloat для получения grade...
	grade, err := getFloat()
	// Если функция вернула ошибку, программа сообщает об этом и завершается.
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failng"
	}

	fmt.Println("A grade of ", grade, " is ", status)
}
