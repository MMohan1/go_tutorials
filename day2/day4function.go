package main

import "fmt"

func sum (x []int)  (float64, [] int){
     var total float64 = 0
     for _ ,value := range x {
         total += float64(value)
     }
     return total, x
     }

func more_values (x ...int )  float64{ // multiple same type values
     var total float64 = 0
     for _ ,value := range x {
         total += float64(value)
     }
     return total
     }


func first() {
  fmt.Println("1st")
}
func second() {
  fmt.Println("2nd")
}

func main () {
     var y = [] int {1,2,3,4,5}
     k, m := sum(y)
     fmt.Println(more_values(y ...))
     fmt.Println(k, m)

     multiply := func (k, l int) int {  // func inside the func
              return k *l
     }

     fmt.Println(multiply(9, 10))
     defer second()
     first()
     // h := map[string]interface{} {}
     h:= map[string]interface{}{"Name": "John", "Age": 35, "locations": []interface{}{"jaipur", "delhi"}}  // example of multitype dict
     // h["name"] = "manmohan"
     // h["age"] = 2
     // h["age"] = {1,23}
     fmt.Println(h)
}