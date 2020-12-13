package main

import "fmt"

func main(){
  var number int
  for true{
    fmt.Printf("Enter the number you want the factors of: ")
    fmt.Scanln(&number)
    factors:= []int{1}
    for i:=2; i<number; i++{
      if number%i==0{
        fmt.Printf("Found a factor %d!\n", i)
        factors=append(factors, i)
      }
    }
    fmt.Printf("The factors %v\n", factors)
  }
}
