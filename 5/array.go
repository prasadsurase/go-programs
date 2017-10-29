package main

import "fmt"

func main() {
  var x [5]int
  var total float64
  x[4] = 100
  fmt.Println(x)

  for i := 0; i< len(x); i++ {
    total += float64(x[i])
  }

  fmt.Println("Total: ", total)

  y := [5]float64{34, 45, 56, 43, 78}
  total = 0.0
  //for i, value := range x {
  for _, value := range y {
    total += float64(value)
  }
  fmt.Println("Total: ", total)
  fmt.Println("=====================================")

  arr := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,}
  min := arr[0]
  for _, value := range arr {
    if(min >= value){
      min = value
    }
  }
  fmt.Println(arr)
  fmt.Println("Shortest: ", min)
}
