package generator

import (
	"html/template"
	"os"
)

const (
	marshallStructKey       = "marshallStruct"
	marshallStructFieldsKey = "marshallStructFields"
)

// Task04 - функция для генерации маршалера структуры в мапу
func MarshallerGenerator(marshallerTemplate string, structName string, inFilePath string, outFilePath string) error {
	return nil
}

// Task03 - структура для yaml конфигурации
// Если нужно расширить yaml файл, тогда эту структуру нужно также расширить
// нужными параметрами, yaml файл и эта структура должны быть идентичными.
type Config struct {
	Name       string
	Port       string
	ReplicaSet int
	ImageName  string
	Tag        string
	EnvPath    string
}

// Task03 - функция генерации конфига
func ConfigGenerate(tmpl string, outFilePath string) error {
	config := &Config{
		Name:       "task_03",
		Port:       "9090",
		ReplicaSet: 13,
		ImageName:  "dfd.png",
		Tag:        "#df8",
		EnvPath:    "module07/internal/generator",
	}
	err := generate(tmpl, outFilePath, config)

	if err != nil {
		return err
	}

	return nil
}

// Основная функция которая должна производить генерацию
func generate(tmpl string, outfilePath string, fields interface{}) error {
	t, err := template.New("configTmpl").Parse(tmpl)
	if err != nil {
		return err
	}

	outFile, err := os.Create(outfilePath)

	if err != nil {
		return err
	}

	defer outFile.Close()

	err = t.Execute(outFile, fields)
	if err != nil {
		return err
	}
	return nil

}
