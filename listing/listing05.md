Что выведет программа? Объяснить вывод программы.
```golang
    package main

    type customError struct {
        msg string
    }
    
    func (e *customError) Error() string {
        return e.msg
    }
    
    func test() *customError {
        {
            // do something
        }
        return nil
    }
    
    func main() {
        var err error
        err = test()
        if err != nil {
            println("error")
            return
        }
        println("ok")
    }
```
Ответ:
`error`
т.к. значение любого интерфейса, является `<nil>` в случае когда И значение И тип являются `<nil>`,
а тип в этом случае равен `*main.customErrorerror`.
