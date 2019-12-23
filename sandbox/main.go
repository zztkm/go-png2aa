package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const AARoot = "./aa"

type AATemplate struct {
	AAs []string
}

func main() {
	AAFiles, err := ioutil.ReadDir(AARoot)
	if err != nil {
		log.Fatal(err)
	}

	tpl := template.Must(template.ParseFiles("./aa.tpl"))
	aas := make([]string, 0)

	for _, AAFile := range AAFiles {
		aaBytes, err := ioutil.ReadFile(filepath.Join(AARoot, AAFile.Name()))
		if err != nil {
			log.Fatal(err)
		}
		aaString := strings.Replace(string(aaBytes), "`", "`+\"+`+\"+`", -1)
		aas = append(aas, aaString)
	}

	aatemplate := AATemplate{aas}
	aaGoSource, err := os.OpenFile("./aa.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	if err := tpl.Execute(aaGoSource, aatemplate); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
}
