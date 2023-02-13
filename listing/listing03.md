Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

    package main
    
    import (
        "fmt"
        "os"
    )
    
    func Foo() error {
        var err *os.PathError = nil
        return err
    }
    
    func main() {
        err := Foo()
        fmt.Println(err)
        fmt.Println(err == nil)
    }
Ответ:
`<nil>`
`false`

В первом случае выводится `<nil>`, так как возвращаемый интерфейсный тип
содержит `os.PathError == nil`.
Во втором случае `false` т.к. значение любого интерфейса, не только error, является `<nil>` в случае когда И значение И тип являются `<nil>`,
а тип в этом случае равен `*fs.PathError`.

Что такое интерфейс?
Интерфейс — абстракция, которая определяет методы которым должен соответствовать тип, чтобы реализовать интерфейс.

Что такое пустой интерфейс?
Интерфейс без методов. Любой тип имплементирует его. Используется для принятия любого типа(как дженерики).