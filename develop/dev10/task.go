package main

/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
    1 Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
    2 Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
    3 При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout

*/

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Client struct { // Стуктура клиента
	host    string
	port    string
	timeout time.Duration
	conn    net.Conn
}

func (client *Client) ConnectionTCP() error { // Подключение к адрессу сервера по адрессу сети до истечения timeout.
	connectStr := fmt.Sprintf("%s:%s", client.host, client.port)
	fmt.Println("Trying", connectStr, "...")
	conn, err := net.DialTimeout("tcp", connectStr, client.timeout)
	if err != nil {
		return err
	}
	fmt.Println("Connected to", connectStr, ".")
	client.conn = conn // интерфейс net.Conn, который можно использовать, как поток вывода или записи
	return nil
}

func (client *Client) sendRequest() { // Получение данных запроса. Используем чтенение через буффер Reader.
	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(client.conn)
	clientRequest, err := reader.ReadString('\n') // Считывем строку методом ReadString(), пока не встрети '\n'.
	if err != nil {                               // Если пользователь вводит 'ctr+D' выходи из программы.
		if err == io.EOF {
			return
		}
		fmt.Println(err)
		return
	}
	_, err = fmt.Fprintf(client.conn, clientRequest)
	for {
		response, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(response)
	}
}

func (client *Client) checkFlag() error {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) <= 1 {
		return errors.New("illegal option")
	}
	client.timeout = time.Duration(5) * time.Second
	order := 0
	if len(argsWithoutProg) == 3 {
		client.parceTimeout(argsWithoutProg)
		order++
	}
	client.host = argsWithoutProg[order]
	order++
	client.port = argsWithoutProg[order]
	return nil
}

func (client *Client) parceTimeout(argsWithoutProg []string) {
	strs := strings.Split(argsWithoutProg[0], "=")
	args := strings.Split(strs[1], "s")
	timeout, _ := strconv.Atoi(args[0])
	client.timeout = time.Duration(timeout) * time.Second
}

func main() {
	client := Client{}
	client.checkFlag()
	err := client.ConnectionTCP()
	if err != nil {
		fmt.Println(err)
		return
	}
	client.sendRequest()
}
