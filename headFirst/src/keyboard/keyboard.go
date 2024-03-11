// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
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
