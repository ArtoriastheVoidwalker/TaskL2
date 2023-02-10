package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Flag struct { // Структура флагов
	A bool
	B bool
	C bool
	c bool
	i bool
	v bool
	F bool
	n bool
}

type Grep struct { // Структура утилиты
	flags    Flag
	pattern  string
	row      int
	cashe    []string
	filename string
	outSet   []string
	outIndex []int
}

func (grep *Grep) CheckFlag() error { // Обработка выбора флага
	argsWithoutProg := os.Args[1:]
	grep.pattern = argsWithoutProg[0]
	for i := 1; i < len(argsWithoutProg)-1; i++ {
		switch argsWithoutProg[i] {
		case "-A":
			grep.flags.A = true
			i++
			err := grep.ParceRow(argsWithoutProg[i])
			if err != nil {
				return errors.New("Invalid argument")
			}
		case "-B":
			grep.flags.B = true
			i++
			err := grep.ParceRow(argsWithoutProg[i])
			if err != nil {
				return errors.New("Invalid argument")
			}
		case "-C":
			grep.flags.C = true
			i++
			err := grep.ParceRow(argsWithoutProg[i])
			if err != nil {
				return errors.New("Invalid argument")
			}
		case "-c":
			grep.flags.c = true
		case "-i":
			grep.flags.i = true
		case "-v":
			grep.flags.v = true
		case "-F":
			grep.flags.F = true
		case "-n":
			grep.flags.n = true
		default:
			return errors.New("Invalid option")
		}
	}
	grep.filename = argsWithoutProg[len(argsWithoutProg)-1]
	return nil
}

func (grep *Grep) ParceRow(str string) error { // Число строк после совпадения
	number, err := strconv.Atoi(str)
	if err != nil {
		return errors.New("Invalid argument")
	}
	if number <= 0 {
		return errors.New("Invalid argument")
	}
	grep.row = number
	return nil
}

func (grep *Grep) ReadFile(filename string) error { // Логика чтения данных из файла в кэш
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grep.cashe = append(grep.cashe, scanner.Text())
	}
	return nil
}

func (grep *Grep) AddFlagA(pos int) { // Добавление строк после совпадения
	for i := pos; i < pos+grep.row+1; i++ {
		if len(grep.outIndex) > 0 && grep.outIndex[len(grep.outIndex)-1] > i {
			continue
		}
		if len(grep.outIndex) > 0 && grep.outIndex[len(grep.outIndex)-1] < i {
			grep.outSet = append(grep.outSet, "--")
			grep.outIndex = append(grep.outIndex, 0)
		}
		grep.outSet = append(grep.outSet, grep.cashe[i])
		grep.outIndex = append(grep.outIndex, i+1)
	}
}

func (grep *Grep) choiceRegex(i int) (bool, error) { // Проверка содержит ли строка какое-либо соответствие паттерну
	check, _ := regexp.MatchString(grep.pattern, grep.cashe[i])
	if grep.flags.i {
		check, _ = regexp.MatchString("(?i)"+grep.pattern, grep.cashe[i])
	}
	if grep.flags.i {
		return !check, nil
	}
	return check, nil
}

func (grep *Grep) choiceRegexFlagF(i int) bool {
	check := grep.cashe[i] == grep.pattern
	if grep.flags.i {
		return !check
	}
	return check
}

func (grep *Grep) AddFlagB(pos int) { // Добавление строк до совпадения
	i := pos - grep.row
	if i < 0 {
		i = 0
	}
	for ; i <= pos; i++ {
		if len(grep.outIndex) > 0 && grep.outIndex[len(grep.outIndex)-1] > i {
			continue
		}
		if len(grep.outIndex) > 0 && grep.outIndex[len(grep.outIndex)-1] < i {
			grep.outSet = append(grep.outSet, "--")
			grep.outIndex = append(grep.outIndex, 0)
		}
		grep.outSet = append(grep.outSet, grep.cashe[i])
		grep.outIndex = append(grep.outIndex, i+1)
	}
}

// Поиск соответствий относительно указаных флагов

func (grep *Grep) WorkFlagA() {
	for i := 0; i < len(grep.cashe); i++ {
		matched, _ := grep.choiceRegex(i)
		if matched {
			grep.AddFlagA(i)
		}
	}
}

func (grep *Grep) WorkFlagB() {
	for i := 0; i < len(grep.cashe); i++ {
		matched, _ := grep.choiceRegex(i)
		if matched {
			grep.AddFlagB(i)
		}
	}
}

func (grep *Grep) WorkFlagC() {
	for i := 0; i < len(grep.cashe); i++ {
		matched, _ := grep.choiceRegex(i)
		if matched {
			grep.AddFlagB(i)
			grep.AddFlagA(i)
		}
	}
}

func (grep *Grep) WorkBasic() {
	for i, str := range grep.cashe {
		matched, _ := grep.choiceRegex(i)
		if matched {
			grep.outSet = append(grep.outSet, str)
			grep.outIndex = append(grep.outIndex, i+1)
		}
	}
}

func (grep *Grep) WorkFlagF() {
	for i, str := range grep.cashe {
		if grep.choiceRegexFlagF(i) {
			grep.outSet = append(grep.outSet, str)
			grep.outIndex = append(grep.outIndex, i+1)
		}
	}
}

func (grep *Grep) Parse() { // Парсинг ключей и опеределение поведения
	if grep.flags.F {
		grep.WorkFlagF()
	} else if grep.flags.A {
		grep.WorkFlagA()
	} else if grep.flags.B {
		grep.WorkFlagB()
	} else if grep.flags.C {
		grep.WorkFlagC()
	} else {
		grep.WorkBasic()
	}
}

func (grep *Grep) Print() { // Вывод результата поиска
	if grep.flags.c {
		fmt.Println(len(grep.outSet))
	} else if grep.flags.n {
		for i, str := range grep.outSet {
			if str == "--" && grep.outIndex[i] == 0 {
				fmt.Println(str)
				continue
			}
			index := fmt.Sprintf("%d:", grep.outIndex[i])
			fmt.Println(index + str)
		}
	} else {
		for _, str := range grep.outSet {
			fmt.Println(str)
		}
	}
}

func main() {
	grep := Grep{}

	err := grep.CheckFlag()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = grep.ReadFile(grep.filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	grep.Parse()
	grep.Print()
}
