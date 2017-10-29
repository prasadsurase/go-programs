package main

import "fmt"

type Person struct {
  Name string
}

type Android struct {
  Person
  Model string
}

func (p *Person) talk() {
  fmt.Println("Hi, my name is", p.Name)
}

func main() {
  p := Person{ Name: "Prasad"}
  p.talk()

  a := new(Android)
  a.Person.talk()
  a.talk()
}
