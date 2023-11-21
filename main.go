package main

import (
  "github.com/snakajima/slashgptGo/manifest"
  "github.com/snakajima/slashgptGo/chatConfig"
  "github.com/snakajima/slashgptGo/chatSession"
  "fmt"
)

func main() {
  m := manifest.New("./sample.json")

  c := chatConfig.New()
  s := chatSession.New(c, m)
  s.Append_message("user", "I am not feeling well")
  res := s.Call_llm()
  fmt.Printf(res)
}
