package main

import (
	"io/ioutil"
	"module07/internal/generator"
	"module07/internal/monitors"
	"os"
)

func Task03() {
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

func Task04() {
	monitorTemplate, err := ioutil.ReadFile("./assets/template/monitor_template.gotmpl")
	if err != nil {
		panic(err)
	}

	monitorTestTemplate, err := ioutil.ReadFile("./assets/template/monitor_test_template.gotmpl")
	if err != nil {
		panic(err)
	}

	if err := monitors.GenerateMonitor(string(monitorTemplate), string(monitorTestTemplate)); err != nil {
		panic(err)
	}
}

func main() {
	// Раскоментировать, когда дойдем до задания номер 3
	// Task03()
	// Раскоментировать, когда дойдем до задания номер 4
	// Task04()
}
