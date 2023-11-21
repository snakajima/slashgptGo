package chatSession

import (
  "github.com/snakajima/slashgptGo/chatConfig"
  "github.com/snakajima/slashgptGo/manifest"
	"context"
	"github.com/sashabaranov/go-openai"
  "fmt"
)

type Session struct {
  Config *chatConfig.Config
  Manifest *manifest.Manifest 
}

func New(config *chatConfig.Config, m *manifest.Manifest) *Session {
  return &Session{config, m}
}

func (session *Session) Call_llm() string {
  systemPrompt := session.Manifest.SystemPrompt()
  messages := []openai.ChatCompletionMessage{
    {
      Role: "system",
      Content: systemPrompt,
    },
    {
      Role: openai.ChatMessageRoleUser,
      Content: "Hello.",
    },
  }
  fmt.Printf(messages[0].Content)
	resp, err := session.Config.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

  return resp.Choices[0].Message.Content
}
