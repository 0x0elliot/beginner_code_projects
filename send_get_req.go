package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main(){
  var site string
  //This would just send a GET request to the site
  fmt.Printf("Give site to send request to [Use http/https]: ")
  fmt.Scanln(&site)
  resp, err:=http.Get(site)
  if err!=nil{
    panic(err)
  }
  data, err:=ioutil.ReadAll(resp.Body)
  defer resp.Body.Close()
  if err!=nil{
    panic(err)
  }
  fmt.Println(string(data))
}
