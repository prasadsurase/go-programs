package main

import "fmt"

func main() {
  var x []int
  fmt.Println("X: ", x)
  y := make([]int, 5, 10)
  fmt.Println("Y: ", y)

  arr := []int{1, 2, 3, 4, 5}
  fmt.Println("Array: ", arr)
  z := arr[0:2]
  fmt.Println("Slice: ", z)
  z = arr[2:4]
  fmt.Println("Slice: ", z)

  slice1 := []int{ 1, 2, 3, 4, 5}
  fmt.Println("Slice1: ", slice1)
  slice2 := append(slice1, 6, 7)
  fmt.Println("Slice2: ", slice2)
  fmt.Println("Slice1: ", slice1)
}
