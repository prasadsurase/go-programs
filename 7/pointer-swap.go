package main

import "fmt"

func swap(x, y *int) {
  temp := *x
  *x = *y
  *y = temp
}

func main() {
  x, y := 1, 2
  fmt.Println("x:", x, ", y:", y)
  swap(&x, &y)
  fmt.Println("x:", x, ", y:", y)

  aPtr, bPtr := new(int), new(int)
  *aPtr, *bPtr = 5, 10
  fmt.Println("a:", *aPtr, ", b:", *bPtr)
  swap(aPtr, bPtr)
  fmt.Println("a:", *aPtr, ", b:", *bPtr)
}
