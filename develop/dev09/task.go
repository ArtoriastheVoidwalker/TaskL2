package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath) // Создание файла
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url) // Получение данных по ссылке
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK { // Проверка получения данных клиентом
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	_, err = io.Copy(out, resp.Body) // Запись страницы в файл
	if err != nil {
		return err
	}
	return nil
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Println("unable to resolve address")
	}
	strs := strings.Split(argsWithoutProg[0], "/")
	err := downloadFile(strs[len(strs)-1], argsWithoutProg[0])
	if err != nil {
		fmt.Println(err)
	}
}
