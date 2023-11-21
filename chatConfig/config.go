package chatConfig

import (
  "os"
)

type Config struct {
	OpenAIKey string
}

func Create() Config {
  key := os.Getenv("OPENAI_API_KEY")
  return Config{key}
}


