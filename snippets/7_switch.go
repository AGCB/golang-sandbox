package main
import (
  "fmt"
  "time"
)
func main() {
  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("today must be a weekend")
  default:
    fmt.Println("thankyou for monday")
  }
}
