package main

import "fmt"

func main() {
  var x map[string]int

  //x["key"] = 10
  fmt.Println(x)

  y := make(map[string]int)
  y["a"] = 10
  y["b"] = 20
  fmt.Println(y)
  fmt.Println(y["b"])
}
