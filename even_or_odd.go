package main

import "fmt"

func main(){
  for true{
    var number int
    fmt.Println("Enter your number: ")
    fmt.Scanln(&number)
    if number%2==0{
      fmt.Println("%d is even", number)
    }else{
      fmt.Println("%d is odd", number)
    }
  }
}
