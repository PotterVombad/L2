package main
 
import (
    "fmt"
    "os"
)
 
func Foo() error {
    var err *os.PathError = nil
    return err
}
 
// программа выведет: <nil>, false
func main() {
    err := Foo()
    // выведет nil т.к. интерфейс содержит указатель на структуру равный nil
    fmt.Println(err)
    // не равно, потому что сам interface error не пустой
    fmt.Println(err == nil)
}
