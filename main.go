package main

import (
  "github.com/konojunya/translate-command/Utils"
  "fmt"
)

func main() {

  text := Utils.Translate("en","ja","hello")

  fmt.Println(text)

}
