package main

import "fmt"

func evenGenerator() func() int {
  g := 0
  return func() (x int) {
    x = g
    g += 2
    return
  }
}

func main() {
  nextEven := evenGenerator()
  fmt.Println(nextEven())
  fmt.Println(nextEven())
  fmt.Println(nextEven())
}
