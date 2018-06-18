package main

import ("fmt")


func main() {
     var x = [] float64  {1,2,3,4}
     // x[0] = 1
     // x[1] = 2
     // x[2] = 3
     // x[3] = 4
     // x[4] = 5

     var total float64 = 0
     for i:=0; i <len(x) ; i++ {
         total += x[i]
     } 
     fmt.Println(total/ float64(len(x)))

     
     x = append(x, 8,9, 10)
     fmt.Println(x)
     for _, value := range x {
         fmt.Println(value)
     }

     k := make(map[string]int)
     k["key"] = 90
     fmt.Println(k)
    
     }