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

  c := chatConfig.New()
  s := chatSession.New(c, m)

  res := chatSession.Call_llm(s)
  fmt.Printf(res)
}
