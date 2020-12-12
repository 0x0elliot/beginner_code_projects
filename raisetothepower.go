package main

import "fmt"
import "math"

func main(){
  var a float64
  var b float64

  for true{
    fmt.Println("Select a")
    fmt.Scanln(&a)
    fmt.Println("Select the value to raise it to (b) : ")
    fmt.Scanln(&b)
    fmt.Println(math.Pow(a, b))
  }
}
