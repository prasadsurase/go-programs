package main

import "fmt"

func factorial(x int) int {
  if x == 0 {
    return 1
  }

  return x * factorial(x - 1)
}

func main() {
  fmt.Println("1: ", factorial(1))
  fmt.Println("2: ", factorial(2))
  fmt.Println("3: ", factorial(3))
  fmt.Println("4: ", factorial(4))
  fmt.Println("5: ", factorial(5))
  fmt.Println("6: ", factorial(6))
  fmt.Println("7: ", factorial(7))
  fmt.Println("8: ", factorial(8))
  fmt.Println("9: ", factorial(9))
}
