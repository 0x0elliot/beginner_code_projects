package main

import (
  "fmt"
  "strings"
)


func main(){
  var name string
  for true{
    fmt.Println("What is your name?")
    fmt.Scanln(&name)
    if strings.ToLower(name)=="rose"{
      fmt.Println("Haha rose u gay brrrrr")
    }else{
      fmt.Println("o okay")
    }
  }
}
