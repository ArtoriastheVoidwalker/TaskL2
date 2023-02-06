package main

import (
	"errors"
	"fmt"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func unpacking(s string) (string, error) {
	sliceRune := []rune(s) // Слайз рун используется чтобы работать с другими раскладками (слайз Байтов не подойдёт т.к. символ может весить от 2-4 байта)
	var result []rune      // Слайз рун для вывода элементов(если использовать строку,
	// то при работе НЕ с аски символы,перед которыми 0, не пропадают а оставляют последовательность битых байт)
	var n int          // Для приведения из rune в int
	var backslash bool // Для обработки escape-последовательностей

	for i, item := range sliceRune {
		if unicode.IsDigit(item) && i == 0 { // Первый элемент является числом
			return "", ErrInvalidString
		}
		if unicode.IsDigit(item) && unicode.IsDigit(sliceRune[i-1]) && sliceRune[i-2] != '\\' { // Обработка случаев с escape-последовательностями
			return "", ErrInvalidString // Когда встречается последовательность из символа и подряд идущих чисел
		}
		if item == '\\' && !backslash { // Является ли символ escape-последовательностью
			backslash = true
			continue
		}
		if backslash && unicode.IsLetter(item) {
			return "", ErrInvalidString
		}
		if backslash { // Добавляем элемент к новой строке, изменяя значение backslash(Завершаем работу с escape-последовательностью)
			result = append(result, item)
			backslash = false
			continue
		}
		if unicode.IsDigit(item) { // Если символ является числом
			n = int(item - '0') // Форматируем из руны в число(числа от 0-9 представленны в виде кодов 48-57,
			// для получения нормального представления необходимо отнять 0)
			if n == 0 {
				result = result[:len(result)-1]
				continue
			}
			for j := 0; j < n-1; j++ { // Добавляем к результату символы,в кол-ве равном следующему за ним числу
				result = append(result, sliceRune[i-1])
			}
			continue
		}
		result = append(result, item)
	}
	return string(result), nil // Перевод слайса рун в строку
}

func main() {
	fmt.Println(unpacking(`кыв0ер\\в4`))
}
