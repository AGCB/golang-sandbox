package main

import ( // wrap multiple imports in parens
         // each on its own line NO comma-separation
  "fmt"
  "math" // math standard lib
)

const thing string = "thing1" // double quotes mandatory.

func main() {
  fmt.Println(thing)
  const thing2 = 5000000
  fmt.Println(math.Sin(thing2)) // -0.9765424686570829
  const thing3 = 3e20 / thing2
  fmt.Println(thing3)           // 6e+13
  fmt.Println(int64(thing3))    // 60000000000000
}
