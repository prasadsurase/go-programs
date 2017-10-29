package main

import "fmt"

func sum(arr []int) int {
  total := 0
  for _, value := range arr {
    total += value
  }

  return total
}

func main() {
  var arr = []int{ 23, 34, 54, 34, 56, 77, 12, 56, 86, 43, 24, 74, 72 }
  total := sum(arr[0:1])
  fmt.Println("Slice: ", arr[0:1])
  fmt.Println("Total: ", total)

  total = sum(arr[0:5])
  fmt.Println("Slice: ", arr[0:5])
  fmt.Println("Total: ", total)

  total = sum(arr[1:8])
  fmt.Println("Slice: ", arr[1:8])
  fmt.Println("Total: ", total)

  total = sum(arr[5:])
  fmt.Println("Slice: ", arr[5:])
  fmt.Println("Total: ", total)

}
