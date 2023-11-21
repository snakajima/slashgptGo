package chatSession

import (
  "github.com/snakajima/slashgptGo/chatConfig"
)

type Session struct {
  Config chatConfig.Config
}

func Create(config chatConfig.Config) Session {
  return Session{config}
}
