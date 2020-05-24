package main
import "fmt"
// feels similar to js bit without paren-wrapping ifCond

func main() {
    // classic
    i:=1
    for i <= 5 { // code in curlys.
      fmt.Println("we can iterate through a for loop with i=", i)
      i = i+1     // iterate in the body.
    }
    //
    //initial/condition/after style
    for j:=40; j<45; j++ {
      fmt.Println("initial/cond/after style and j value of ",j)
    }
    //
    // // loop infinitely with explicit break.
    for {
      fmt.Println("could have been forever") // dont slip up on single quotes
                                             // burnt few cycles debugging
      break // i appreciate a removal of semicolons from jsLife
    }
    //
    // skip iteration with continue
    for z:=0; z<=10; z++ {
      if z%2 == 0 { // if even number... equality with double equals
        continue;
      }
      fmt.Println("testing continue with z of",z)
    }
    //

}
