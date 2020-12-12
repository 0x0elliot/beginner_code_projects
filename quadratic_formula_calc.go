package main

import "fmt"
import "math"

func main(){
  var a,b,c,x1,x2,d float64

  for true{

    fmt.Println("Welcome to my Roots Of Quadratic Formula calculator.")
    fmt.Printf("Enter the value of coefficient of x^2: ")
    fmt.Scanln(&a)
    if a==0{
      fmt.Println("The value of a can't be zero!")
      return 
    }
    fmt.Printf("Enter the value of coefficient of x: ")
    fmt.Scanln(&b)
    fmt.Printf("Enter the value of the constant c: ")
    fmt.Scanln(&c)
    //Calculating the roots
    d=b*b - 4*a*c
    if d<0{
      fmt.Println("Roots are imaginary!")
      return
    }
    x1= (-b + math.Pow(d, 0.5))/(2*a)
    x2= (-b - math.Pow(d, 0.5))/(2*a)

    fmt.Sprintf("The root of the equation %fx^2 + %fx + %f:", a, b,c)
    fmt.Println(x1)
    fmt.Println(x2)
}
}
