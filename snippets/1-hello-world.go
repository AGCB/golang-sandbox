package main //

import "fmt" // format is a standard lib

func main() { // similar to JS. "main" used twice here raises Qs??
              // can main use another function defined in the same file? -- YES... defined after OR before works fine.
              // if so, does it matter if the function is defined before or after? -- 1st attempt failed on this.

  fmt.Println("hello_World.3") // fmt.Sentance case()
                             // "" denote strings

  myPrinter()  // call a func
}

func myPrinter() {
  fmt.Println("myPrinter function")
}

// run in terminal with ``` go run filename ```
