// This tool generates the example data by loading the json file from the
// openslides repo.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

const (
	exampleDataURL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/openslides4-dev/docs/example-data.json"
	outFile        = "example-data.json.go"
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

	f, err := os.Create(outFile)
	if err != nil {
		log.Fatalf("Can not create output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Can not close file: %v", err)
		}
	}()

	if err := writeFile(f, data); err != nil {
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
	var d map[string][]map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&d); err != nil {
		return nil, fmt.Errorf("decoding full file: %w", err)
	}

	data := make(map[string]string)
	for collection, v := range d {
		for i, element := range v {
			var id int
			if err := json.Unmarshal(element["id"], &id); err != nil {
				return nil, fmt.Errorf("decoding id of %dth element in collection %s", i, collection)
			}

			for k, v := range element {
				data[fmt.Sprintf("%s/%d/%s", collection, id, k)] = string(v)
			}
		}
	}
	return data, nil
}

const tpl = `// Code generated with example-data.json DO NOT EDIT.
package main

import "encoding/json"

var exampleData = map[string]json.RawMessage{
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
		"Escape": string(escape),
		"Data":   eData,
	}

	if err := t.Execute(replacer{w}, data); err != nil {
		return fmt.Errorf("writing template: %w", err)
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

func intsToStr(ints []int) string {
	var s string
	for _, i := range ints {
		s += fmt.Sprintf("%d,", i)
	}
	return strings.TrimSuffix(s, ",")
}
