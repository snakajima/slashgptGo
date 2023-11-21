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
  fmt.Println(m.SystemPrompt())

  c := chatConfig.New()
  s := chatSession.New(c, m)

  res := s.Call_llm()
  fmt.Printf(res)
}
