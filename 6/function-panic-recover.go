package main

import "fmt"

func main() {
  defer func() {
    str := recover()
    fmt.Println("Error:", str)
  }()
  panic("something went wrong again")

  panic("something went wrong")
  str := recover()
  fmt.Println("Error: ", str)

}
