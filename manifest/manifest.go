package manifest

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "strings"
)

type Manifest struct {
	Title string `json:"title"`
  Prompt []string `json:"prompt"`
}

func New(path string) *Manifest {
  fileBytes, err := ioutil.ReadFile(path)
  if err != nil {
      fmt.Println(err)
      return &Manifest{}
  }
  var m Manifest
  err = json.Unmarshal(fileBytes, &m)
  if err != nil {
      fmt.Println(err)
      return &Manifest{}
  }

  return &m
}

func (m *Manifest) SystemPrompt() string {
  return strings.Join(m.Prompt, "\n")
}