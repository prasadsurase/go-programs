package main

import "fmt"

var outer string = "outer hello world"

func main() {
  fmt.Println("Outer: ", outer)
  inner := "inner hello world"
  fmt.Println("Inner: ", inner)
  {
    inner = "scoped hello world"
    fmt.Println("Scoped: ", inner)
  }
  fmt.Println("Inner: ", inner)
}
