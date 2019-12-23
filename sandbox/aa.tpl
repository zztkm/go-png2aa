package main

import (
    "fmt"
    "math/rand"
    "time"
)

var aas = []string{
    {{range $AA := .AAs}}
`
{{$AA}}
`,{{end}}
}

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println(aas[rand.Intn(len(aas))])
}
