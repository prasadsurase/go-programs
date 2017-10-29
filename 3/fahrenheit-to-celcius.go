package main

import "fmt"

func main() {
  fmt.Print("Fahrenheit: ")
  var fah int
  fmt.Scanf("%d", &fah)
  cel := ((fah - 32) * 5/9.0)
  fmt.Println("Celsius: ", cel)
}
