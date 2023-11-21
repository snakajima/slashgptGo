package chatConfig

import (
  "os"
  "github.com/sashabaranov/go-openai"
)

type Config struct {
	OpenAIKey string
  Client* openai.Client
}

func New() Config {
  key := os.Getenv("OPENAI_API_KEY")
  client := openai.NewClient(key)
  return Config{key, client}
}


