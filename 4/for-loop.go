package main

import "fmt"

func main() {
  i := 1
  for i <= 5 {
    fmt.Println("I: ", i)
    i += 1
  }

  for i := 1; i <= 5; i++ {
    fmt.Println("I: ", i)
  }
}
