package main

import "fmt"

func main(){
  // This is my first time using arrays in go

  var myarr[4]string
  //defining elements.
  myarr[0]="Hi"
  myarr[1]="I"
  myarr[2]="am"
  myarr[3]="Elliot"
  //Weird flex but ok. range=last index+1 which is weird but i mean go is weird but lovely
  fmt.Println(myarr)
}
