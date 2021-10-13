// This tool generates the example data by loading the json file from the
// openslides repo.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

const (
	exampleDataURL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/master/docs/example-data.json"
	packageName    = "models"
)

func main() {
	e, err := loadExampleData()
	if err != nil {
		log.Fatalf("Can not load example data: %v", err)
	}
	defer e.Close()

	data, err := decode(e)
	if err != nil {
		log.Fatalf("Can not decode example data: %v", err)
	}

	if err := writeFile(os.Stdout, data); err != nil {
		log.Fatalf("Can not write result: %v", err)
	}
}

func loadExampleData() (io.ReadCloser, error) {
	r, err := http.Get(exampleDataURL)
	if err != nil {
		return nil, fmt.Errorf("request defition: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request returned status %s", r.Status)
	}
	return r.Body, nil
}

func decode(r io.Reader) (map[string]string, error) {
	var decoded map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&decoded); err != nil {
		return nil, fmt.Errorf("decoding full file: %w", err)
	}

	data := make(map[string]string)
	for collection, collectionData := range decoded {
		if collection == "_migration_index" {
			continue
		}

		var decodedCollection map[string]map[string]json.RawMessage
		if err := json.Unmarshal(collectionData, &decodedCollection); err != nil {
			return nil, fmt.Errorf("decoding collection %s: %w", collection, err)
		}

		for id, w := range decodedCollection {
			for field, value := range w {
				data[fmt.Sprintf("%s/%s/%s", collection, id, field)] = string(value)
			}
		}
	}
	return data, nil
}

const tpl = `// Code generated with example-data.json DO NOT EDIT.
package {{ .pkg }}

import "encoding/json"

// ExampleData is a generated value from the OpenSlides example data.
//
// It is a map from key (fqfield) to the value encoded to json.
var ExampleData = map[string]json.RawMessage{
	{{- range $key, $value := .Data}}
	"{{$key}}": []byte({{$.Escape}}{{$value}}{{$.Escape}}),
	{{- end}}
}
`

func writeFile(w io.Writer, eData map[string]string) error {
	t := template.New("t")
	t, err := t.Parse(tpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	data := map[string]interface{}{
		"pkg":    packageName,
		"Escape": string(escape),
		"Data":   eData,
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(replacer{buf}, data); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}

	formated, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("formating code: %w", err)
	}

	if _, err := w.Write(formated); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
}

// The output needs the backtick (`) to work. But a backtick can not be used in
// a backtick-string. Therefore we use another byte and replace this byte with a
// backtick afterwards.
const escape byte = 1

type replacer struct {
	w io.Writer
}

func (r replacer) Write(p []byte) (n int, err error) {
	for i, b := range p {
		if b == escape {
			p[i] = '`'
		}
	}
	return r.w.Write(p)
}
