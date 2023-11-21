package main

import (
  "github.com/snakajima/slashgptGo/manifest"
  "github.com/snakajima/slashgptGo/chatConfig"
  "github.com/snakajima/slashgptGo/chatSession"
  "fmt"
)

func main() {
  m := manifest.New("./sample.json")
  fmt.Println(m.Title)    
  fmt.Println(m.Prompt)

  c := chatConfig.Create()
  s := chatSession.Create(c)
  fmt.Println(s.Config.OpenAIKey)
}
