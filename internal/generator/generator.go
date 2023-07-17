package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"os"
	"strings"
)

const (
	marshallStructKey       = "marshallStruct"
	marshallStructFieldsKey = "marshallStructFields"
)

// Task04 - функция для генерации маршалера структуры в мапу
func MarshallerGenerator(marshallerTemplate string, structName string, inFilePath string, outFilePath string) error {

	type Field struct {
		Name  string
		Value string
	}

	type templateData struct {
		Package    string
		StructName string
		Fields     []Field
	}

	templData := &templateData{}
	templField := &Field{}

	fileSet := token.NewFileSet()

	node, err := parser.ParseFile(fileSet, inFilePath, nil, parser.ParseComments)

	templData.StructName = structName
	templData.Package = node.Name.Name
	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if strucT, ok := typeSpec.Type.(*ast.StructType); ok {
						fieldsTemp := make([]Field, 0, len(strucT.Fields.List))
						for _, field := range strucT.Fields.List {
							var name string
							nameValue := field.Names[0].Name
							if field.Tag != nil {
								name = strings.Split(field.Tag.Value, "\"")[1]
							} else {
								name = field.Names[0].Name
							}
							templField.Name = name
							templField.Value = nameValue
							fieldsTemp = append(fieldsTemp, *templField)
							templData.Fields = fieldsTemp

						}
					}
				}
			}
		}
	}

	if err != nil {
		return err
	}

	if err := generate(marshallerTemplate, outFilePath, templData); err != nil {
		return err
	}

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
		fmt.Println(err)
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
