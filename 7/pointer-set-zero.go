package main

import "fmt"

func zero(x int){
  x = 0
}

func setZero(xPtr *int){
  *xPtr = 0
}

func setOne(xPtr *int) {
  *xPtr = 1
}

func main() {
  x := 5
  zero(x)
  fmt.Println(x)

  setZero(&x)
  fmt.Println(x)

  xPtr := new(int)

  fmt.Println(*xPtr)
  setOne(xPtr)
  fmt.Println(*xPtr)
}
