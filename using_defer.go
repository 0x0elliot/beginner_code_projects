package main

import "fmt"

func main(){
  defer fmt.Println("This comes up after main returns/Does it's job and ends even though this is written before the next print sentence")

  fmt.Println("So this is why")
}
