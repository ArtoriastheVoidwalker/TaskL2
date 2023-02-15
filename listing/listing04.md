Что выведет программа? Объяснить вывод программы.
```golang
    package main
    
    func main() {
        ch := make(chan int)
        go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
        }()

        for n := range ch {
            println(n)
        }
    }
```
Ответ:
    `0
    1
    2
    3
    4
    5
    6
    7
    8
    9
    fatal error: all goroutines are asleep - deadlock!`

`range` будет читать данные из канала,пока тот открыт. Из-за того что канал не закрывается и оттуда нечего 
читать получаем ошибку. Её можно исправить просто закрыв канал:
```golang
    func main() {
        ch := make(chan int)
        go func() {
            defer close(ch)
            for i := 0; i < 10; i++ {
                ch <- i
            }
        }()
        for n := range ch {
            println(n)
        }
    }
```