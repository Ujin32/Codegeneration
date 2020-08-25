package main

import (
	"io/ioutil"
	"module07/internal/generator"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("./assets/template/config_template.yml")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("./assets/template/config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := generator.Generate(string(data), f); err != nil {
		panic(err)
	}
}
