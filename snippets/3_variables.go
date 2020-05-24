package main
import "fmt"
func main() {
  var a = "initial" // same as js
  fmt.Println(a)
  var b,c int = 3,4 // define multiple vars..
                    // notice here we typed to integer
  fmt.Println(b,c)

  var d = false // why did we need to strongly type int but not bool or string?
  var e = true  // both bools are lowercase
  fmt.Println(d,e)

  var f int // "zero-valued" .. for int is 0
  fmt.Println(f)

  g:="apple"
  fmt.Println(g)
}
