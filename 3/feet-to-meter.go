package main

import "fmt"

func main() {
  const meter float64 = 0.3048
  var input int
  fmt.Print("Enter number of feet: ")
  fmt.Scanf("%d", &input)
  meters := float64(input) * meter
  fmt.Println("Meters: ", meters)
}
