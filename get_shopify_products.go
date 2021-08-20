package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "strings"
  "strconv"
)

func main() {

  var site string

  fmt.Println("Shopify Site URL:")

  fmt.Scanln(&site)

  for i := 1;  i < 3; i++ {

    url := site+"/products.json?page="+strconv.Itoa(i)

    method := "GET"

    client := &http.Client {
    }
    req, err := http.NewRequest(method, url, nil)

    if err != nil {
      fmt.Println(err)
      return
    }
    res, err := client.Do(req)
    if err != nil {
      fmt.Println(err)
      return
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
      fmt.Println(err)
      return
    }
    data:=string(body)
    if(strings.Count(data,"")>16){
      ioutil.WriteFile(strconv.Itoa(i)+".json", []byte(data), 0644)
    }
  }
  fmt.Println("Done.")
}
