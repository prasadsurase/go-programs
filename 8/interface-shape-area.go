package main

import "fmt"
import "math"

type Shape interface {
  area() float64
  perimeter() float64
}

type Circle struct {
  radius float64
}

type Rectangle struct {
  length, breadth float64
}

type Square struct {
  side float64
}

func (c *Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func (r *Rectangle) area() float64 {
  return r.length * r.breadth
}

func (s *Square) area() float64 {
  return s.side * s.side
}

func (c *Circle) perimeter() float64 {
  return 2 * math.Pi * c.radius
}

func (r *Rectangle) perimeter() float64 {
  return 2 * (r.length + r.breadth)
}

func (s *Square) perimeter() float64 {
  return 4 * s.side
}

func main(){
  r := Rectangle{ length: 12.5, breadth: 6.5 }
  c := Circle{ radius: 6.5 }
  s := Square{ side: 6.5 }

  fmt.Println("Area of Rectangle:", r.area() )
  fmt.Println("Area of Circle:", c.area() )
  fmt.Println("Area of Square:", s.area() )

  fmt.Println("Perimeter of Rectangle:", r.perimeter() )
  fmt.Println("Perimeter of Circle:", c.perimeter() )
  fmt.Println("Perimeter of Square:", s.perimeter() )
}
