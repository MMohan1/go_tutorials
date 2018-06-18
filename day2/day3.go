package main

import ("fmt"
       )

func main () {
     x  := []int {1, 2, 3, 4, 5}
     fmt.Println(x)
     for _, value := range x {
     fmt.Println(value)
     }

     var i int = 0
     for i < len(x) {  // while implimenation in go
         fmt.Println(x[i])
         i += 1
     }

y := [] int {1,2,3}
k := make([] int, 7)
z := append(y, 4,5)
// z[4] = 90
l := append(k, 90)
l[7] = 9
fmt.Println(z, l)


 a :=  make(map[string] int)
 a["manmohan"] = 1
 a["manmohan"] = 8
 fmt.Println(a)
 c, d := a["kgmanmohan"]
 fmt.Println(c, d)
 }