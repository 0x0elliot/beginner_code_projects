package main

import "fmt"
import "os"

func main(){
  var number int
  for true{
    fmt.Printf("Enter the number to check whether it's a prime number or not: ")
    fmt.Scanln(&number)
    for i:=2; i<number; i++{
      if number%i==0{
        fmt.Println("The number is NOT prime!")
        os.Exit(1)
      }
    }
    fmt.Println("The number is prime!")
  }
}
