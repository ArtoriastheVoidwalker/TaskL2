package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func readFileToStringsLines(dir string) (result [][]string) { // Для чтения данных из файла и преобразования в двумерный слайс
	file, err := os.Open(dir) // Открываем файл для чтения
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) // Создание сканера

	for scanner.Scan() {
		var words []string
		words = strings.Split(scanner.Text(), " ") // Разделение строк на слова
		result = append(result, words)             // Запись строк в слайс
	}

	return result
}

func writeToFile(data [][]string, dir string) { // Для записи данных в файл
	file, err := os.Create(dir) // Создание файла
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, len(data)) // Преобразование данных в текст
	for i, datum := range data {
		str := strings.Join(datum, " ")
		lines[i] = str
	}
	_, err = file.WriteString(strings.Join(lines, "\n")) // Запись в файл
	if err != nil {
		panic(err)
	}
}

func getMonth(month string) time.Time {
	if t, err := time.Parse("Jan", month); err == nil {
		return t
	}
	if t, err := time.Parse("January", month); err == nil {
		return t
	}
	if t, err := time.Parse("1", month); err == nil {
		return t
	}
	if t, err := time.Parse("01", month); err == nil {
		return t
	}
	return time.Time{}
}

func getLen(str []string) int {
	var result int
	result = len(str) - 1
	for _, s := range str {
		result += len(s)
	}
	return result
}

func getDataElem(data [][]string, i, k int) string { // Возвращает элемент с индексом i
	if k < len(data[i]) {
		return data[i][k]
	}
	return ""
}

func getDataElemSlice(data [][]string, i, k int) []string { // Получение элементов строки i начиная с индекса k
	if k < len(data[i]) {
		return data[i][k:]
	}
	return []string{}
}

func main() {
	var n, r, u, m, b, c, h bool // Для ключей
	var k int
	var data [][]string              // Для сохранения считанной инф из файла
	var sortFunc func(i, j int) bool // Для выбора метод сортировки

	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&m, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&h, "h", false, "сортировать по числовому значению с учётом суффиксов")
	flag.BoolVar(&c, "c", false, "проверять отсортированы ли данные")

	// считываем ключи параметров сортировки
	flag.IntVar(&k, "k", 0, "указание индекса колонки для сортировки")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")

	flag.Parse()

	// считываем путь к сортируемому файлу
	input := flag.Arg(0)
	output := flag.Arg(1)

	if input == "" || output == "" { // Проверка ввода имён файлов
		panic("Input or output file name cannot be empty.")
	}

	if k < 1 { //
		k = 0
	} else {
		k--
	}

	data = readFileToStringsLines(input) // Чтение содержимого файла

	switch true {
	case n: // Сортировка по числовому значению

		sortFunc = func(i, j int) bool {
			a, _ := strconv.ParseFloat(getDataElem(data, i, k), 64)
			b, _ := strconv.ParseFloat(getDataElem(data, j, k), 64)
			if r {
				return a > b
			}
			return a < b
		}
	case m: // Сортировка по месяцам
		sortFunc = func(i, j int) bool {
			if r {
				return getMonth(getDataElem(data, j, k)).Before(getMonth(getDataElem(data, i, k)))
			}
			return getMonth(getDataElem(data, i, k)).Before(getMonth(getDataElem(data, j, k)))
		}
	case h: // Сортировка по количеству символов в строке
		sortFunc = func(i, j int) bool {
			if r {
				return getLen(data[i][k:]) > getLen(data[j][k:])
			}
			return getLen(getDataElemSlice(data, i, k)) < getLen(getDataElemSlice(data, j, k))
		}
	default: // Сортировка строки
		sortFunc = func(i, j int) bool {
			if r {
				return getDataElem(data, i, k) > getDataElem(data, j, k)
			}
			return getDataElem(data, i, k) < getDataElem(data, j, k)
		}
	}
	if c { // Проверка введен ли ключ проверки сортировки, если да - проверяем файл на упорядоченность
		isSorted := sort.SliceIsSorted(data, sortFunc)
		fmt.Println("Sorted?", isSorted)
		return
	}
	sort.Slice(data, sortFunc) // Сортировка согласно заданному ключу
	writeToFile(data, output)  // Запись отсортированных данных в указанный файл
}
