// This tool generates the list of related fields in the file def.go.
// To call it, just call "go generate ./..." in the root folder of the repository
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

const defURL = "https://raw.githubusercontent.com/normanjaeckel/OpenSlides/modelsToYML/docs/models.yml"

func main() {
	r, err := loadDefition()
	if err != nil {
		log.Fatalf("Can not load models defition: %v", err)
	}
	defer r.Close()

	data, err := parse(r)
	if err != nil {
		log.Fatalf("Can not parse model definition: %v", err)
	}

	if err := writeFile(os.Stdout, data); err != nil {
		log.Fatalf("Can not write result: %v", err)
	}
}

func loadDefition() (io.ReadCloser, error) {
	r, err := http.Get(defURL)
	if err != nil {
		return nil, fmt.Errorf("request defition: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request returned status %s", r.Status)
	}
	return r.Body, nil
}

// parse returns all relation-list and generic-relation-list fields and where
// they point to.
func parse(r io.Reader) (map[string]string, error) {
	var inData map[string]map[string]mValue
	if err := yaml.NewDecoder(r).Decode(&inData); err != nil {
		return nil, fmt.Errorf("decoding models.yml: %w", err)
	}

	outData := make(map[string]string)

	for model, fields := range inData {
		for field, value := range fields {
			if value.Type == "template" {
				value = value.template.Fields
			}

			if value.Type == "relation-list" || value.Type == "generic-relation-list" {
				outData[fmt.Sprintf("%s/%s", model, field)] = value.relation.toCollection()
				continue
			}

		}
	}
	return outData, nil
}

const tpl = `// Code generated with models.txt DO NOT EDIT.
package restrict

var relationLists = map[string]string{
	{{- range $key, $value := .Def}}
	"{{$key}}": "{{$value}}",
	{{- end}}
}
`

func writeFile(w io.Writer, rlist map[string]string) error {
	t := template.New("t")
	t, err := t.Parse(tpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	data := map[string]interface{}{
		"Def": rlist,
	}

	if err := t.Execute(w, data); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}
	return nil
}
