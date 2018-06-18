package main

import ("fmt")
/*

func main() {
     // var a float64 = 9
     // var b int
     var (
     a float64 = 9
     b int
     )
     fmt.Println("please provide a int number ==>")
     fmt.Scanf("%d",  &b)
     
     fmt.Println("the squrt of the number is", math.Sqrt(a))
     fmt.Println("the tuble of the input is  ",  b)
     }

*/


func main(){
     for  i:=0; i<10; i++ {
     if i %2 == 0{
        fmt.Println(i, "divided by 2")
        } else if i%3 == 0 {
        fmt.Println(i, "divided by 3")
        }
     }

     // for  i:=0; i<10; i++ {
     // switch i {
     //        case i % 2 == 0: {
     //             fmt.Println(i, "divided by 2")
     //             }
     //      case i % 3 == 0: {
     //           fmt.Println(i, "divided by 3")
     //           }
     //     default: fmt.Println(i, "divided by Unknown")
     // }
     // }


}