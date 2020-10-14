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

	"github.com/OpenSlides/openslides-modelsvalidate/models"
)

const defURL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/openslides4-dev/docs/models.yml"

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
	inData, err := models.Unmarshal(r)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling models.yml: %w", err)
	}

	outData := make(map[string]string)
	for modelName, model := range inData {
		for attrName, attr := range model.Attributes {
			r := attr.Relation()

			if r == nil {
				continue
			}

			to := r.ToCollection()
			if len(to) != 1 {
				to[0] = "*"
			}
			outData[fmt.Sprintf("%s/%s", modelName, attrName)] = to[0]
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
