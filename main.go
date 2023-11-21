package main

import (
  "github.com/snakajima/slashgptGo/manifest"
  "github.com/snakajima/slashgptGo/chatConfig"
  "github.com/snakajima/slashgptGo/chatSession"
  "fmt"
  "bufio"
  "os"
	"github.com/sashabaranov/go-openai"
)

func main() {
  m := manifest.New("./sample.json")

  c := chatConfig.New()
  s := chatSession.New(c, m)
  fmt.Print("You: ")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  message := scanner.Text()
  s.Append_message(openai.ChatMessageRoleUser, message)
  res := s.Call_llm()
  fmt.Println("GPT:", res)
}
