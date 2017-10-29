package main

import "fmt"

func average(xs []float64)(int, float64) {
  total := 0.0
  for _, value := range xs {
    total += value
  }
  return len(xs), total/float64(len(xs))
}

func main() {
  xs := []float64{98, 93, 77, 82, 83, 50}
  fmt.Println("Array: ", xs)
  size, avg := average(xs)
  fmt.Println("Length: ", size)
  fmt.Println("Average: ", avg)
}
