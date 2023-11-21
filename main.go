package main

import (
  "github.com/snakajima/slashgptGo/manifest"
  "github.com/snakajima/slashgptGo/chatConfig"
  "github.com/snakajima/slashgptGo/chatSession"
  "fmt"
	"context"
	"github.com/sashabaranov/go-openai"
)

func main() {
  m := manifest.New("./sample.json")
  fmt.Println(m.Title)    
  fmt.Println(m.Prompt)

  c := chatConfig.New()
  s := chatSession.New(c)
  fmt.Println(s.Config.OpenAIKey)

	client := openai.NewClient(s.Config.OpenAIKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Tell me about Starwars movies!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
