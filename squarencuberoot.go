package main

import "fmt"
import "math"

func main(){
  var i int
  i=1
  fmt.Println("Welcome to my square root calculator.")
  for i<2{
    var num,root float64
    fmt.Println("Enter your number: ")
    fmt.Scanln(&num)
    root=math.Pow(num, 0.5)
    fmt.Println(root)
  }
}
