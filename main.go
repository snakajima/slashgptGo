package main

import (
  "github.com/snakajima/slashgptGo/manifest"
  "fmt"
)

func main() {
  m := manifest.Load("./sample.json")
  fmt.Println(m.Title)    
  fmt.Println(m.Prompt)
}
