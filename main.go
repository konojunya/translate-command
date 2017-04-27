package main

import(
  "net/http"
  "fmt"
  "net/url"
  "strings"
)

func main() {
  // text := "こんばんわ"
  // translate_type := "ja/en"

  // data := {
  //   "client_id": ,
  //   "client_secret": ,
  //   "scope": "http://api.microsofttranslator.com",
  //   'grant_type': 'client_credentials',
  // }

  values := url.Values{}
  values.Add("host","datamarket.accesscontrol.windows.net")
  values.Add("path","/v2/OAuth2-13")

  req, err := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
  if err != nil {
    fmt.Println(err)
    return
  }

  client := new(http.Client)
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }

  defer resp.Body.Close()

  fmt.Println(resp.Body)

}