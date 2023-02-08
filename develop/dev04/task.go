package main

import (
	"fmt"
	"sort"
	"strings"
)

func searchAnagram(arr []string) map[string][]string { // Основная функция поиска анограммы
	mapAnagram := make(map[string][]string)
	iterationArray(arr, mapAnagram)
	deleteSingle(mapAnagram)
	sortMap(mapAnagram)
	return mapAnagram
}

func iterationArray(arr []string, mapAnagram map[string][]string) {
	for _, str := range arr {
		check := true
		strLower := strings.ToLower(str)
		for key, val := range mapAnagram {
			if determinateAnagram(key, strLower) {
				if !binarySearch(val, strLower) { // Проверка наличия строки среди значений, добавление в случае отсутствия
					mapAnagram[key] = append(val, strLower)
				}
				check = false
				break
			}
		}
		if check {
			mapAnagram[strLower] = append(mapAnagram[strLower], strLower) // Получаем необработанный словарь из анограмм
		}
	}
}

func determinateAnagram(str1 string, str2 string) bool { // Выявление анограмм

	if len(str1) != len(str2) { // Если длинна слов разная они не будут анаграммами
		return false
	}
	setOne := make(map[rune]int)
	for _, ch := range str1 { // Подсчёт символов первого слова
		setOne[ch] += 1
	}
	setTwo := make(map[rune]int)
	for _, ch := range str2 { // Подсчёт символов второго слова
		setTwo[ch] += 1
	}

	for key, val := range setOne { // Если пары ключ/значения слов не равны-слова не анаграммы
		if val != setTwo[key] {
			return false
		}
	}
	return true
}

func sortMap(mapAnagram map[string][]string) { // Сортировка словаря
	for _, val := range mapAnagram {
		sort.Slice(val, func(i, j int) bool { return val[i] < val[j] })
	}
}

func deleteSingle(mapAnagram map[string][]string) { // Удаление из словаря ключей без анограмм
	for key, val := range mapAnagram {
		if len(val) == 1 { // Если значений меньше двух-убираем из словаря
			delete(mapAnagram, key)
		}
	}
}

func binarySearch(arr []string, pattern string) bool { // Бинарный поиск слова в слайсе
	max, min := len(arr), 0
	for i := 0; i < len(arr); i++ {
		index := (max - min) / 2
		if arr[index] == pattern {
			return true
		} else if arr[index] < pattern {
			min = index
		} else {
			max = index
		}
	}
	return false
}

func main() {
	letters := []string{"слиток",
		"автобус",
		"пятка",
		"Столик",
		"Столик",
		"тяпка"}
	m := searchAnagram(letters)
	fmt.Println(m)
}
