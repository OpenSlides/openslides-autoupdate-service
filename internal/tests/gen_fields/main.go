// This tool generates the list fields of every collection.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	models "github.com/OpenSlides/openslides-models-to-go"
)

// TODO unkomment after this is merged: https://github.com/OpenSlides/OpenSlides/pull/5802
// const defURL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/openslides4-dev/docs/models.yml"
const defURL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/7847a90f22e04f563618e4bc7dab894e517bd9e3/docs/models.yml"

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

// parse returns all fqfields ordered by collection.
func parse(r io.Reader) (map[string][]string, error) {
	inData, err := models.Unmarshal(r)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling models.yml: %w", err)
	}

	outData := make(map[string][]string)
	for modelName, model := range inData {
		for fieldName := range model.Fields {
			outData[modelName] = append(outData[modelName], fieldName)
		}
	}

	return outData, nil
}

const tpl = `// Code generated with models.txt DO NOT EDIT.
package tests

var collectionFields = map[string][]string{
	{{- range $key, $value := .Def}}
	"{{$key}}": { {{range $v := $value}} "{{$v}}", {{end}} },
	{{- end}}
}
`

func writeFile(w io.Writer, rlist map[string][]string) error {
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
