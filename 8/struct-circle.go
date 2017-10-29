package main

import "fmt"
import "math"

type Circle struct {
  x, y, r float64
}

func circleArea(r float64) float64 {
  return math.Pi * r * r
}

func areaCircle(c *Circle) float64{
  return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

func main() {
  c1 := Circle{x: 0, y: 0, r: 5}

  fmt.Println("Area:", circleArea(c1.r))
  fmt.Println("Area:", areaCircle(&c1))
  fmt.Println("Area:", c1.area())

}
