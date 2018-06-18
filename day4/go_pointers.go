package main

import "fmt"

func main () {
     a := 10
     b := &a
     *b = 11
     fmt.Println("the poniter memory refeance is", b)
     fmt.Println("the poniter valuee is", *b)
     fmt.Println("the poniter value of a is", a)

}