Что выведет программа? Объяснить вывод программы. 
Рассказать про внутреннее устройство слайсов и что происходит при передаче их в качестве аргументов функции.
```golang
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```
Ответ:
`[3 2 3]`

Слайс - это структура данных с переменной длиной ссылающаяся на массив, также имеющая параметры длины и ёмкости. 
В данном случае при добавлении элемента в слайс не хватает емкости
(количество элементов, которых можно добавить в слайс без реаллокации), 
и будет создан новый массив с увлеченной вдвое capacity, 
элементы из старого массива будут скопированы в новый, 
слайс теперь будет ссылаться уже на новый массив.