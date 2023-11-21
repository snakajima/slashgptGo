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
  Messages []openai.ChatCompletionMessage;
}

func New(config *chatConfig.Config, m *manifest.Manifest) *Session {
  messages := []openai.ChatCompletionMessage{
    {
      Role: openai.ChatMessageRoleSystem,
      Content: m.SystemPrompt(),
    },
  }
  return &Session{config, m, messages}
}

func (session *Session) Call_llm() string {
	resp, err := session.Config.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: session.Messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

  return resp.Choices[0].Message.Content
}

func (session *Session) Append_message(role string, message string) {
  session.Messages = append(session.Messages, openai.ChatCompletionMessage{Role:role, Content:message})
}