package main

import "fmt"
import "math"

type Rectangle struct{
  x1, x2, y1, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := y1 - x1
  b := y2 - x2
  return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x2, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y2)
  return l * w
}

func main() {
  r := Rectangle{0, 0, 10, 10}
  fmt.Println(r.area())
}
