package monitors

import (
	"fmt"
	"os"
	"text/template"
)

type MonitorGeneratorConfig struct {
	Monitor               string
	MonitorType           string
	MonitorConfig         string
	ConfigInterfaceFields []ConfigMonitorFields
}

type ConfigMonitorFields struct {
	Name      string
	Type      string
	TestValue interface{}
}

var monitorFields = map[string]MonitorGeneratorConfig{
	"simple_monitor": {
		Monitor:       "SimpleMonitor",
		MonitorType:   "monitor_type_simple_monitor",
		MonitorConfig: "SimpleMonitorConfig",
		ConfigInterfaceFields: []ConfigMonitorFields{
			{Name: "Field1", Type: "string", TestValue: `"simple"`},
			{Name: "Field2", Type: "float64", TestValue: 1.2345},
			{Name: "Field3", Type: "int", TestValue: 234},
		},
	},
}

func GenerateMonitor(monitorTemplate string, monitorTestTemplate string) error {
	for fileName, fields := range monitorFields {
		if err := generateMonitor(monitorTemplate, fileName, fields); err != nil {
			return err
		}

		if err := generateMonitorTest(monitorTestTemplate, fileName, fields); err != nil {
			return err
		}
	}

	return nil
}

func generateMonitor(monitorTemplate string, fileName string, fields MonitorGeneratorConfig) error {
	f, err := os.Create(fmt.Sprintf("./internal/monitors/%s.go", fileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := f.Truncate(0); err != nil {
		return err
	}

	t, err := template.New("monitor").Parse(monitorTemplate)
	if err != nil {
		return err
	}

	return t.Execute(f, fields)
}

func generateMonitorTest(monitorTestTemplate string, fileName string, fields MonitorGeneratorConfig) error {
	f, err := os.Create(fmt.Sprintf("./internal/monitors/%s_test.go", fileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := f.Truncate(0); err != nil {
		return err
	}

	t, err := template.New("monitor_test").Parse(monitorTestTemplate)
	if err != nil {
		return err
	}

	return t.Execute(f, fields)
}
