package main

import (
    "github.com/snakajima/slashgptGo/greet"
    "github.com/snakajima/slashgptGo/manifest"
    "fmt"
)

func main() {
    m := manifest.Create("Hello")
    fmt.Println(m.Title)    
    fmt.Println(greet.Message())
}
