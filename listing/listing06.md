package main
 
import (
  "fmt"
)
// выведет 3, 2, 3
func main() {
  var s = []string{"1", "2", "3"}
  modifySlice(s)
  fmt.Println(s)
}

// при добавлении элемента у нас создастся новый слайс, вследствии чего изменен будет только первый элемент
func modifySlice(i []string) {
  i[0] = "3"
  i = append(i, "4")
  i[1] = "5"
  i = append(i, "6")
}