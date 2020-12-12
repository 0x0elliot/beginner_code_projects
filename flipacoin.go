package main

import ("fmt"
"math/rand")

func main(){

  for true{
    fmt.Println("Press enter to flip a coin.")
    fmt.Scanln()
    rando:=rand.Intn(2)
    if rando==0{
      fmt.Println("You got heads!")
    }else{
      fmt.Println("You got tails!")
    }
  }
}
