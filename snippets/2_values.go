package main

import "fmt"

func printValues() {
  fmt.Println("go" + "lang") // string concatenation
  fmt.Println(42, 42+2) // integers and math
  fmt.Println(4.0/2.1) // floats
  fmt.Println(true && false) // AND
  fmt.Println(true || false) // OR
}

func main() {
  printValues()

}
