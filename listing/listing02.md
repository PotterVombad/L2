package main
 
import (
    "fmt"
)
 
func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return
}
 
 
func anotherTest() int {
    var x int
    defer func() {
        x++
    }()
    x = 1
    return x
}
 
//программа выведет: 2, 1
func main() {
    // в первом случае мы работаем с именованной переменной в возвращаемом значении, функция переданная в defer работает с этой переменной, следовательно изменяет ее и она становится равна 2
    fmt.Println(test())
    // во втором случае у нас анонимное возвращаемое значение, т.е. мы передали в возвращаемую анонимную переменную значение, но работать с ней уже не можем, т.к. обращаться не к кому, вследствие чего мы просто изменяем локальную переменную, а не возвращаемое значение
    fmt.Println(anotherTest())
}
