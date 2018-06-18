/*
package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}
func rectangleArea(x1, y1, x2, y2 float64) float64 {
  l := distance(x1, y1, x1, y2)
  w := distance(x1, y1, x2, y1)
  return l * w
}
func circleArea(x, y, r float64) float64 {
  return math.Pi * r*r
}
func main() {
  var rx1, ry1 float64 = 0, 0
  var rx2, ry2 float64 = 10, 10
  var cx, cy, cr float64 = 0, 0, 5

  fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
  fmt.Println(circleArea(cx, cy, cr))
}
*/


package main


import ("fmt"; "reflect"; "math")

type Circle struct {
     x, y, r float64
 }

func (c Circle) circleArea() float64 {
     return math.Pi * c.r*c.r
}

func (c *Circle) modifyCircleRedious()  {
     c.r = 10
}

func main() {
     c := Circle{x: 10.0, y:20.0, r:9}
     fmt.Println(reflect.TypeOf(c.x))
     c.modifyCircleRedious()
     fmt.Println(c.circleArea())

}
