package main

import (
  "github.com/snakajima/slashgptGo/manifest"
  "github.com/snakajima/slashgptGo/config"
  "fmt"
)

func main() {
  m := manifest.Load("./sample.json")
  fmt.Println(m.Title)    
  fmt.Println(m.Prompt)

  c := config.Create()
  fmt.Println(c.OpenAIKey)    
}
