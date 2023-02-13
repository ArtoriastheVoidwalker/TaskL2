package main

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

*/

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type Shell struct { // Структура утилиты
	dir string
}

func parceArgs() []string { // Реализация парсинга аргументов
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return strings.Split(scanner.Text(), " | ")
	}
	return nil
}

// Реализация логики команд

func (shell *Shell) cd(command string) {
	args := command[len("cd")+1:]
	os.Chdir(args)
	newDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	shell.dir = newDir + " %"
}

func (shell *Shell) pwd() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
}

func (shell *Shell) echo(command string) {
	args := command[len("echo")+1:]
	fmt.Println(args)
}

func (shell *Shell) kill(command string) {
	args := command[len("echo")+1:]
	pid, _ := strconv.Atoi(args)
	err := syscall.Kill(pid, 9)
	if err != nil {
		fmt.Println(err)
	}
}

func (shell *Shell) ps(str string) {
	out, err := exec.Command(str).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func main() {
	shell := Shell{
		dir: "%",
	}
	for {
		fmt.Printf("%s ", shell.dir)
		commands := parceArgs()
		for _, command := range commands {
			if strings.Index(command, "cd") == 0 {
				shell.cd(command)
			} else if strings.Index(command, "pwd") == 0 {
				shell.pwd()
			} else if strings.Index(command, "echo") == 0 {
				shell.echo(command)
			} else if strings.Index(command, "kill") == 0 {
				shell.kill(command)
			} else if strings.Index(command, "ps") == 0 {
				shell.ps(command)
			} else if strings.Index(command, "\\quit") == 0 ||
				strings.Index(command, "\\q") == 0 {
				return
			} else {
				fmt.Println("command not found:", command)
			}
		}
	}
}
