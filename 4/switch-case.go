package main

import "fmt"

func main() {
  var i int
  fmt.Println("Enter 1 to 3: ")
  fmt.Scanf("%d", &i)
  switch i {
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
    case 3: fmt.Println("Three")
    default: fmt.Println("Duffer")
  }

  for i:= 1; i<= 100; i++ {
    //switch i {
    switch {
      case ((i % 3 == 0) && (i % 5 == 0)): fmt.Println(i, ": Fizz Buzz")
      case (i % 3 == 0): fmt.Println(i, ": Fizz")
      case (i % 5 == 0): fmt.Println(i, ": Buzz")
      default: fmt.Println(i)
    }
  }
}
