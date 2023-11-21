package main

import (
    "github.com/snakajima/slashgptGo/greet"
    "github.com/snakajima/slashgptGo/manifest"
    "fmt"
)

func main() {
    m := manifest.Create("Hello")
    fmt.Println(m.Title)    
    m2 := manifest.Load("./sample.json")
    fmt.Println(m2.Title)    
    fmt.Println(greet.Message())
}
