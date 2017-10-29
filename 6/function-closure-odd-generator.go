package main

import "fmt"

func oddGenerator() func() int{
  x := 1
  return func() (i int){
    i = x
    x += 2
    return
  }
}

func main(){
  nextOdd := oddGenerator()
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
}
