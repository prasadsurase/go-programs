package main

import "fmt"

func average(args ...float64)(int, float64) {
  fmt.Println("Args: ", args)
  total := 0.0
  for _, value := range args {
    total += value
  }
  return len(args), total/float64(len(args))
}

func main() {
  //xs := []float64{98, 93, 77, 82, 83, 50}
  size, avg := average(98, 34, 65, 23, 45, 76, 34, 45)
  fmt.Println("Length: ", size)
  fmt.Println("Average: ", avg)
}
