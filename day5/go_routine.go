package main

import ("fmt")

func run_loop(max_limit int) {
     for i:=0; i< max_limit; i++ {
         fmt.Println(i)
         }

}



func main() {
     for i:=0;i <10000; i++{
     fmt.Println("calling function ==>", i)
     go run_loop(i)
     }

}