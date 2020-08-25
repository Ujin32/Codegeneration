package generator

import (
	"io"
	"text/template"
)

type Config struct {
	Name       string
	Port       string
	ReplicaSet int
	ImageName  string
	Tag        string
	EnvPath    string
}

func Generate(tmp string, w io.Writer) error {
	cfg := Config{
		Name:       "MyApp",
		Port:       "8080",
		ReplicaSet: 3,
		ImageName:  "test_image",
		Tag:        "v1",
		EnvPath:    "./.env",
	}

	t, err := template.New("config").Parse(tmp)
	if err != nil {
		return err
	}

	if err := t.Execute(w, cfg); err != nil {
		return err
	}

	return nil
}
