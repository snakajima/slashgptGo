package chatSession

import (
  "github.com/snakajima/slashgptGo/chatConfig"
  "github.com/snakajima/slashgptGo/manifest"
)

type Session struct {
  Config chatConfig.Config
  Manifest manifest.Manifest 
}

func New(config chatConfig.Config, m manifest.Manifest) Session {
  return Session{config, m}
}

func call_llm(session Session) string {
  return "abc"
}
