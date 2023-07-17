package main

import (
	"fmt"
	"io/ioutil"
	"module07/internal/generator"
	"module07/internal/monitors"
)

func Task03() {
	yamlConfiTemplate, err := ioutil.ReadFile("./assets/template/config_template.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(yamlConfiTemplate))
	_ = generator.ConfigGenerate(string(yamlConfiTemplate), "./config.yml")
}

func Task04() {
	marshallerTemplate, err := ioutil.ReadFile("D:\\golang\\rebrain_basic\\module07\\assets\\template\\marshaller.gotmpl")
	//./assets/template/marshaller.gotmpl
	if err != nil {
		panic(err)
	}

	_ = generator.MarshallerGenerator(
		string(marshallerTemplate),
		"Config",
		"D:\\golang\\rebrain_basic\\module07\\internal\\config\\config.go",
		"D:\\golang\\rebrain_basic\\module07\\internal\\config\\codegen_marshaller.go",
	)
}

// Заполнить логикой, когда дойдем до задания номер 5
func Task05() {
	_ = monitors.NewSimpleMonitor()
}

func main() {
	//Task03()
	//Раскоментировать, когда дойдем до задания номер 4
	Task04()
	// Раскоментировать, когда дойдем до задания номер 5
	// Task05()

	fmt.Println("Hello, from 07 module")
}
