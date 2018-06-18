
package main

import ("fmt";
       "time"
)


func sender (c chan string) {
     for i:=0; ;i++ {
         msg := "ping " + string(i)
         fmt.Println(msg)
         c <- msg
     }

}

func ponger(c chan string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func reciver(c chan string) {
     for {
        // msg := <- c
        fmt.Println( "&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
        time.Sleep(time.Second * 10)
}
var input string
  fmt.Scanln(&input)
}


func main(){
     var c chan string = make (chan string, 1)
     fmt.Println("I am Here")
     go sender(c)
     go ponger(c)
     go reciver(c)
     var input string
     fmt.Scanln(&input)    
     

}