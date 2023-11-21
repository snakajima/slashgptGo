package main

import (
  "github.com/snakajima/slashgptGo/manifest"
  "github.com/snakajima/slashgptGo/chatConfig"
  "fmt"
)

func main() {
  m := manifest.Load("./sample.json")
  fmt.Println(m.Title)    
  fmt.Println(m.Prompt)

  c := chatConfig.Create()
  fmt.Println(c.OpenAIKey)    
}
