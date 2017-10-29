package main

import "fmt"

func findMax(args ...int) int {
  max := args[0]
  for _, value := range args {
    if(value >= max) {
      max = value
    }
  }
  return max
}

func main() {
  fmt.Println(findMax(1, 2))
  fmt.Println(findMax(23, 45, 98, 23, 345, 234, 12))
  fmt.Println(findMax(1, 1))
}
