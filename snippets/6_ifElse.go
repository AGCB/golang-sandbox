package main
import "fmt"

func main() {
  //
  // basic if/else
  if 7%2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }
  //
  // else is optional
  //
  // you can insert statements before conditionals and
  // they will be available to all branches.
  if num:= -87; num < 0 {
    fmt.Println("the number must be negative", num)
  } else if num >= 10 {
    fmt.Println("the number must have 2 digits", num)
  } else {
    fmt.Println("the number must be single digit", num)
  }

}
